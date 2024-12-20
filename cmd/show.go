package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/sharon-xa/memepulation/file"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show stuff based on the user desire",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && !cmd.Flags().Changed("content") && !cmd.Flags().Changed("repeated") &&
			!cmd.Flags().Changed("repeated-number") {
			cmd.Help()
			return
		}
		f = file.ReadFile(filePath)

		content, err := cmd.Flags().GetBool("content")
		if err != nil {
			log.Fatalf("Error retrieving 'content' flag: %v", err)
		}

		repeated, err := cmd.Flags().GetBool("repeated")
		if err != nil {
			log.Fatalf("Error retrieving 'repeated' flag: %v", err)
		}

		repeatedNumber, err := cmd.Flags().GetBool("repeated-number")
		if err != nil {
			log.Fatalf("Error retrieving 'repeated-number' flag: %v", err)
		}

		// Handle flags
		if content {
			f.PrintFileContent()
		}

		if repeated {
			f.PrintRepeatedLines()
		}

		if repeatedNumber {
			f.PrintRepeatedLinesNumber()
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().BoolP("content", "c", false, "Show file content")
	showCmd.Flags().BoolP("repeated", "r", false, "Show all repeated lines in the file")
	showCmd.Flags().
		BoolP("repeated-number", "n", false, "Show the number of repeated lines in the file")
}
