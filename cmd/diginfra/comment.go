package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/spf13/cobra"

	"github.com/diginfra/diginfra/internal/apiclient"
	"github.com/diginfra/diginfra/internal/logging"

	"github.com/diginfra/diginfra/internal/clierror"
	"github.com/diginfra/diginfra/internal/config"
	"github.com/diginfra/diginfra/internal/output"
)

type CommentOutput struct {
	Body           string
	HasDiff        bool
	ValidAt        *time.Time
	AddRunResponse apiclient.AddRunResponse
}

var (
	validCommentOutputFormats = []string{
		"json",
	}
)

func commentCmd(ctx *config.RunContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "comment",
		Short: "Post an Diginfra comment to GitHub, GitLab, Azure Repos or Bitbucket",
		Long:  "Post an Diginfra comment to GitHub, GitLab, Azure Repos or Bitbucket",
		Example: `  Update the Diginfra comment on a GitHub pull request:

      diginfra comment github --repo my-org/my-repo --pull-request 3 --path diginfra.json --behavior update --github-token $GITHUB_TOKEN

  Delete old Diginfra comments and post a new comment to a GitLab commit:

      diginfra comment gitlab --repo my-org/my-repo --commit 2ca7182 --path diginfra.json --behavior delete-and-new --gitlab-token $GITLAB_TOKEN

  Post a new comment to an Azure Repos pull request:

      diginfra comment azure-repos --repo-url https://dev.azure.com/my-org/my-project/_git/my-repo --pull-request 3 --path diginfra.json --behavior new --azure-access-token $AZURE_ACCESS_TOKEN`,
		ValidArgs: []string{"--", "-"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmds := []*cobra.Command{commentGitHubCmd(ctx), commentGitLabCmd(ctx), commentAzureReposCmd(ctx), commentBitbucketCmd(ctx)}
	for _, subCmd := range cmds {
		subCmd.RunE = checkAPIKeyIsValid(ctx, subCmd.RunE)

		subCmd.Flags().StringArray("policy-path", nil, "Path to Diginfra policy files, glob patterns need quotes (experimental)")
		subCmd.Flags().Bool("show-all-projects", false, "Show all projects in the table of the comment output")
		subCmd.Flags().Bool("show-changed", false, "Show only projects in the table that have code changes")
		subCmd.Flags().Bool("show-skipped", true, "List unsupported resources")
		_ = subCmd.Flags().MarkHidden("show-changed")
		subCmd.Flags().Bool("skip-no-diff", false, "Skip posting comment if there are no resource changes. Only applies to update, hide-and-new, and delete-and-new behaviors")
		_ = subCmd.Flags().MarkHidden("skip-no-diff")
		subCmd.Flags().String("comment-path", "", "Path to comment content file (experimental)")
		_ = subCmd.Flags().MarkHidden("comment-path")
	}

	cmd.AddCommand(cmds...)

	return cmd
}

func buildCommentOutput(cmd *cobra.Command, ctx *config.RunContext, paths []string, mdOpts output.MarkdownOptions) (*CommentOutput, error) {
	inputs, err := output.LoadPaths(paths)
	if err != nil {
		return nil, err
	}

	combined, err := output.Combine(inputs)
	if errors.As(err, &clierror.WarningError{}) {
		logging.Logger.Warn().Msg(err.Error())
	} else if err != nil {
		return nil, err
	}

	combined.IsCIRun = ctx.IsCIRun()

	var commentData string
	var governanceFailures output.GovernanceFailures
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	var result apiclient.AddRunResponse
	if ctx.IsCloudUploadEnabled() && !dryRun {
		if ctx.Config.IsSelfHosted() {
			logging.Logger.Warn().Msg("Diginfra Cloud is part of Diginfra's hosted services. Contact hello@diginfra.khulnasoft.com for help.")
		} else {
			combined.Metadata.DiginfraCommand = "comment"
			result = shareCombinedRun(ctx, combined, inputs)
			combined.RunID, combined.ShareURL, combined.CloudURL, governanceFailures = result.RunID, result.ShareURL, result.CloudURL, result.GovernanceFailures
			commentData = result.CommentMarkdown
		}
	}

	var out *CommentOutput

	commentPath, _ := cmd.Flags().GetString("comment-path")
	if commentPath != "" {
		commentData, err = output.LoadCommentData(commentPath)
		if err != nil {
			return nil, fmt.Errorf("Error loading %s used by --comment-path flag. %s", commentPath, err)
		}
	}

	if commentData != "" {
		// the full comment markdown has been received from the API addRun or loaded from the comment-path file,
		// so use that instead of building the output using the output.ToMarkdown templates.
		out = &CommentOutput{
			Body:           commentData,
			HasDiff:        combined.HasDiff(),
			ValidAt:        &combined.TimeGenerated,
			AddRunResponse: result,
		}
	}

	var policyChecks output.PolicyCheck
	policyPaths, _ := cmd.Flags().GetStringArray("policy-path")
	if len(policyPaths) > 0 {
		policyChecks, err = queryPolicy(policyPaths, combined)
		if err != nil {
			return nil, err
		}

		ctx.ContextValues.SetValue("passedPolicyCount", len(policyChecks.Passed))
		ctx.ContextValues.SetValue("failedPolicyCount", len(policyChecks.Failures))
	}

	if out == nil {
		opts := output.Options{
			DashboardEndpoint: ctx.Config.DashboardEndpoint,
			NoColor:           ctx.Config.NoColor,
			PolicyOutput:      output.NewPolicyOutput(policyChecks),
		}
		opts.ShowAllProjects, _ = cmd.Flags().GetBool("show-all-projects")
		opts.ShowOnlyChanges, _ = cmd.Flags().GetBool("show-changed")
		opts.ShowSkipped, _ = cmd.Flags().GetBool("show-skipped")

		md, err := output.ToMarkdown(combined, opts, mdOpts)
		if err != nil {
			return nil, err
		}

		b := md.Msg
		ctx.ContextValues.SetValue("truncated", md.OriginalMsgSize != md.RuneLen)
		ctx.ContextValues.SetValue("originalLength", md.OriginalMsgSize)

		out = &CommentOutput{
			Body:           string(b),
			HasDiff:        combined.HasDiff(),
			ValidAt:        &combined.TimeGenerated,
			AddRunResponse: result,
		}
	}

	if policyChecks.HasFailed() {
		return out, policyChecks.Failures
	}
	if len(governanceFailures) > 0 {
		return out, governanceFailures
	}

	return out, nil
}

type PRNumber int

func (p *PRNumber) Set(value string) error {
	if value == "" {
		return nil
	}

	v, err := strconv.Atoi(value)
	*p = PRNumber(v)

	if err != nil {
		return errors.New("must be integer")
	}

	return nil
}

func (p *PRNumber) String() string {
	return fmt.Sprintf("%d", *p)
}

func (p *PRNumber) Type() string {
	return "int"
}

func queryPolicy(policyPaths []string, input output.Root) (output.PolicyCheck, error) {
	checks := output.PolicyCheck{
		Enabled: true,
	}

	inputValue, err := ast.InterfaceToValue(input)
	if err != nil {
		return checks, fmt.Errorf("Unable to process Diginfra output into Rego input: %s", err.Error())
	}

	ctx := context.Background()
	r := rego.New(
		rego.Query("data.diginfra.deny"),
		rego.ParsedInput(inputValue),
		rego.Load(policyPaths, func(abspath string, info os.FileInfo, depth int) bool {
			return false
		}),
	)
	pq, err := r.PrepareForEval(ctx)
	if err != nil {
		return checks, fmt.Errorf("Unable to query provided policies: %s", err.Error())
	}

	res, err := pq.Eval(ctx)
	if err != nil {
		return checks, err
	}

	if len(res) == 0 {
		return checks, fmt.Errorf("The provided polices returned no valid data.diginfra.deny rules. Please check that the policies are formatted correctly.")
	}

	for _, e := range res[0].Expressions {
		switch v := e.Value.(type) {
		case map[string]interface{}:
			readPolicyOut(v, &checks)
		case []interface{}:
			for _, ii := range v {
				if m, ok := ii.(map[string]interface{}); ok {
					readPolicyOut(m, &checks)
				}
			}
		}
	}

	return checks, nil
}

func readPolicyOut(v map[string]interface{}, checks *output.PolicyCheck) {
	if _, ok := v["msg"]; !ok {
		checks.Failures = append(checks.Failures, "Policy rule invalid as it did not contain {msg: string} property in output object. Please edit rule output object.")
		return
	}
	msg := v["msg"].(string)

	if _, ok := v["failed"]; !ok {
		checks.Failures = append(checks.Failures, fmt.Sprintf("Policy rule: [%s] did not contain {failed: bool} output property. Please edit rule output object.", msg))
		return
	}

	failed, _ := v["failed"].(bool)

	if failed {
		checks.Failures = append(checks.Failures, msg)
		return
	}

	checks.Passed = append(checks.Passed, msg)
}

func isErrorUnhandled(err error) bool {
	if err == nil {
		return false
	}

	switch err.(type) {
	case output.PolicyCheckFailures, output.GovernanceFailures:
		return false
	}

	return true
}
