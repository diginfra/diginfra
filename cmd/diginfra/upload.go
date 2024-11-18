package main

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"

	"github.com/spf13/cobra"

	"github.com/diginfra/diginfra/internal/apiclient"
	"github.com/diginfra/diginfra/internal/config"
	"github.com/diginfra/diginfra/internal/logging"
	"github.com/diginfra/diginfra/internal/output"
	"github.com/diginfra/diginfra/internal/ui"
)

var (
	validUploadOutputFormats = []string{
		"json",
	}
)

func uploadCmd(ctx *config.RunContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload an Diginfra JSON file to Diginfra Cloud",
		Long: `Upload an Diginfra JSON file to Diginfra Cloud. This is useful if you
do not use 'diginfra comment' and instead want to define run metadata,
such as pull request URL or title, and upload the results manually.

See https://diginfra.khulnasoft.com/docs/features/cli_commands/#upload-runs`,
		Example: `  Upload an Diginfra JSON file:
      export DIGINFRA_VCS_PULL_REQUEST_URL=http://github.com/myorg...
      export DIGINFRA_VCS_PULL_REQUEST_TITLE="My PR title"
      # ... other env vars here

      diginfra diff --path plan.json --format json --out-file diginfra.json

      diginfra upload --path diginfra.json`,
		ValidArgs: []string{"--", "-"},
		RunE: checkAPIKeyIsValid(ctx, func(cmd *cobra.Command, args []string) error {
			var err error

			format, _ := cmd.Flags().GetString("format")
			format = strings.ToLower(format)
			if format != "" && !contains(validUploadOutputFormats, format) {
				ui.PrintUsage(cmd)
				return fmt.Errorf("--format only supports %s", strings.Join(validOutputFormats, ", "))
			}

			if ctx.Config.IsSelfHosted() {
				return fmt.Errorf("Diginfra Cloud is part of Diginfra's hosted services. Contact hello@diginfra.khulnasoft.com for help.")
			}

			path, _ := cmd.Flags().GetString("path")

			root, err := output.Load(path)
			if err != nil {
				return fmt.Errorf("could not load input file %s err: %w", path, err)
			}

			dashboardClient := apiclient.NewDashboardAPIClient(ctx)
			result, err := dashboardClient.AddRun(ctx, root)
			if err != nil {
				return fmt.Errorf("failed to upload to Diginfra Cloud: %w", err)
			}

			if format == "json" {
				b, err := jsoniter.MarshalIndent(result, "", "  ")
				if err != nil {
					return fmt.Errorf("failed to marshal result: %w", err)
				}
				cmd.Print(string(b))
			} else if result.ShareURL != "" {
				cmd.Println("Share this cost estimate: ", ui.LinkString(result.ShareURL))
			}

			pricingClient := apiclient.GetPricingAPIClient(ctx)
			err = pricingClient.AddEvent("diginfra-upload", ctx.EventEnv())
			if err != nil {
				logging.Logger.Warn().Err(err).Msg("could not report `diginfra-upload` event")
			}

			return nil
		}),
	}

	cmd.Flags().String("path", "p", "Path to Diginfra JSON file.")
	cmd.Flags().String("format", "", "Output format: json")

	_ = cmd.MarkFlagRequired("path")
	_ = cmd.MarkFlagFilename("path", "json")
	return cmd
}
