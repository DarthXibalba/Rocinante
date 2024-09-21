package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script",
	Long: `To load completion scripts:

Bash:

$ source <(rocinante completion bash)

Zsh:

$ source <(rocinante completion zsh)

Fish:

$ rocinante completion fish | source

PowerShell:

$ rocinante completion powershell | Out-String | Invoke-Expression
`,
	Hidden: true, // This hides the completion command from the help output
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "bash":
				rootCmd.GenBashCompletion(os.Stdout)
			case "zsh":
				rootCmd.GenZshCompletion(os.Stdout)
			case "fish":
				rootCmd.GenFishCompletion(os.Stdout, true)
			case "powershell":
				rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
			default:
				cmd.Help()
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	// Add the completion command to the root command
	rootCmd.AddCommand(completionCmd)
}
