package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/diginfra/diginfra/internal/apiclient"
	"github.com/diginfra/diginfra/internal/config"
)

func authCmd(ctx *config.RunContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Get a free API key, or log in to your existing account",
		Long:  "Get a free API key, or log in to your existing account",
		Example: fmt.Sprintf(`  Get a free API key, or log in to your existing account:

      diginfra auth login

      You can also log in at %v

  Manually set the API key that your CLI should use. The API key can be retrieved from your account:

      diginfra configure set api_key MY_API_KEY

  Regenerate your API key:

      Log in at %v > select your organization > Settings`, ctx.Config.DashboardEndpoint, ctx.Config.DashboardEndpoint),
		ValidArgs: []string{"--", "-"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmds := []*cobra.Command{authLoginCmd(ctx)}
	cmd.AddCommand(cmds...)

	return cmd
}

func authLoginCmd(ctx *config.RunContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate the CLI with your Diginfra account",
		Long:  "Authenticate the CLI with your Diginfra account",
		Example: fmt.Sprintf(`  Get a free API key, or log in to your existing account:

      diginfra auth login

      You can also log in at %v

  Manually set the API key that your CLI should use. The API key can be retrieved from your account:

      diginfra configure set api_key MY_API_KEY`, ctx.Config.DashboardEndpoint),
		ValidArgs: []string{"--", "-"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			cmd.Println("We're redirecting you to our log in page, please complete that,\nand return here to continue using Diginfra.")

			auth := apiclient.AuthClient{Host: ctx.Config.DashboardEndpoint}
			apiKey, info, err := auth.Login(ctx.ContextValues.Values())
			if err != nil {
				return err
			}

			if info != "" {
				cmd.Println(info)
				return nil
			}

			ctx.Config.Credentials.APIKey = apiKey
			ctx.Config.Credentials.PricingAPIEndpoint = ctx.Config.PricingAPIEndpoint

			err = ctx.Config.Credentials.Save()
			if err != nil {
				return err
			}

			fmt.Printf("The API key was saved to %s\n", config.CredentialsFilePath())
			cmd.Println("\nYour account has been authenticated. Run Diginfra on your Terraform project by running:")
			cmd.Printf("\n  diginfra breakdown --path=.\n\n")

			return nil
		},
	}

	return cmd
}
