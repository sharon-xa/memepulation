package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/sharon-xa/memepulation/file"
)

var (
	f        *file.File
	filePath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "memepulation",
	Short: "This program helps you perform operations on text files.",
	Long:  `This CLI allows you to perform various operations on text files, such as showing content or calculating statistics.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "File to be manipulated")
	rootCmd.MarkPersistentFlagRequired("file")

	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `To load completions:

Bash:

$ source <(memepulation completion bash)

# To load completions for each session, execute once:
Linux:
  $ memepulation completion bash > /etc/bash_completion.d/memepulation
MacOS:
  $ memepulation completion bash > /usr/local/etc/bash_completion.d/memepulation

Zsh:

$ source <(memepulation completion zsh)

# To load completions for each session, execute once:
$ memepulation completion zsh > "${fpath[1]}/_memepulation"

Fish:

$ memepulation completion fish | source

# To load completions for each session, execute once:
$ memepulation completion fish > ~/.config/fish/completions/memepulation.fish

Powershell:

PS> memepulation completion powershell | Out-String | Invoke-Expression

# To load completions for every new session, run:
PS> memepulation completion powershell > memepulation.ps1
# and source this file from your PowerShell profile.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		switch args[0] {
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			rootCmd.GenPowerShellCompletion(os.Stdout)
		default:
			cmd.Help()
		}
	},
}
