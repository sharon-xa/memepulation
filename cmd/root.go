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
}
