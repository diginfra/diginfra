package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func completionCmd() *cobra.Command {
	completionCmd := &cobra.Command{
		Use:   "completion --shell [bash | zsh | fish | powershell]",
		Short: "Generate shell completion script",
		Long: `To load completions:
	
	Bash:
	
		$ source <(diginfra completion --shell bash)
	
		# To load completions for each session, execute once:
		# Linux:
		$ diginfra completion --shell bash > /etc/bash_completion.d/diginfra
		# macOS:
		$ diginfra completion --shell bash > /usr/local/etc/bash_completion.d/diginfra
	
	Zsh:
	
		# If shell completion is not already enabled in your environment,
		# you will need to enable it.  You can execute the following once:
	
		$ echo "autoload -U compinit; compinit" >> ~/.zshrc
	
		# To load completions for each session, execute once:
		$ diginfra completion --shell zsh > "${fpath[1]}/_diginfra"
	
		# You will need to start a new shell for this setup to take effect.
	
	fish:
	
		$ diginfra completion --shell fish | source
	
		# To load completions for each session, execute once:
		$ diginfra completion --shell fish > ~/.config/fish/completions/diginfra.fish
	
	PowerShell:
	
		PS> diginfra completion --shell powershell | Out-String | Invoke-Expression
	
		# To load completions for every new session, run:
		PS> diginfra completion --shell powershell > diginfra.ps1
		# and source this file from your PowerShell profile.
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if hasShellFlag := cmd.Flags().Changed("shell"); hasShellFlag {
				shell, err := cmd.Flags().GetString("shell")
				if err != nil {
					return err
				}

				switch shell {
				case "bash":
					_ = cmd.Root().GenBashCompletion(cmd.OutOrStdout())
				case "zsh":
					_ = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
				case "fish":
					_ = cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
				case "powershell":
					_ = cmd.Root().GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
				default:
					return fmt.Errorf("unsupported shell type: %q", shell)
				}
			}

			return nil
		},
	}

	completionCmd.Flags().String("shell", "", "supported shell formats: bash, zsh, fish, powershell")
	_ = completionCmd.MarkFlagRequired("shell")

	_ = completionCmd.RegisterFlagCompletionFunc("shell", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"bash\tCompletions for bash",
				"zsh\tCompletions for zsh",
				"fish\tCompletions for fish",
				"powershell\tCompletions for powershell"},
			cobra.ShellCompDirectiveDefault
	})

	return completionCmd
}
