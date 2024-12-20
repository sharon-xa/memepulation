package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/sharon-xa/memepulation/file"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Display file statistics (total, unique, empty, and duplicate lines)",
	Long: `The stats command analyzes a file and provides statistics such as:
- Total lines
- Unique lines
- Duplicate lines
- Empty lines
- File size and bytes`,
	Run: func(cmd *cobra.Command, args []string) {
		f = file.ReadFile(filePath)

		total, err := cmd.Flags().GetBool("total")
		if err != nil {
			log.Fatalf("Error retrieving 'total' flag: %v", err)
		}

		unique, err := cmd.Flags().GetBool("unique")
		if err != nil {
			log.Fatalf("Error retrieving 'unique' flag: %v", err)
		}

		duplicate, err := cmd.Flags().GetBool("duplicate")
		if err != nil {
			log.Fatalf("Error retrieving 'duplicate' flag: %v", err)
		}

		empty, err := cmd.Flags().GetBool("empty")
		if err != nil {
			log.Fatalf("Error retrieving 'empty' flag: %v", err)
		}

		size, err := cmd.Flags().GetBool("size")
		if err != nil {
			log.Fatalf("Error retrieving 'size' flag: %v", err)
		}

		bytes, err := cmd.Flags().GetBool("bytes")
		if err != nil {
			log.Fatalf("Error retrieving 'bytes' flag: %v", err)
		}

		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatalf("Error retrieving 'all' flag: %v", err)
		}

		if all {
			f.PrintAllStats()
			return
		}

		// Implement logic for each flag
		if total {
			fmt.Printf(
				"Total lines in file '%s': %d\n",
				f.FileName,
				f.GetTotalLinesNum(),
			)
		}

		if unique {
			fmt.Printf(
				"Unique lines in file '%s': %d\n",
				f.FileName,
				f.GetUniqueLines(),
			)
		}

		if duplicate {
			fmt.Printf(
				"Duplicate lines in file '%s': %d\n",
				f.FileName,
				f.GetDuplicateLines(),
			)
		}

		if empty {
			fmt.Printf("Empty lines in file '%s': %d\n", f.FileName, f.GetEmptyLines())
		}

		if size {
			fmt.Printf("Size of file '%s': %s\n", f.FileName, f.GetFileSize())
		}

		if bytes {
			fmt.Printf("Number of bytes in file '%s': %d\n", f.FileName, f.GetFileBytes())
		}

		if !bytes && !size && !empty && !duplicate && !unique && !total && !all {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	statsCmd.Flags().BoolP("total", "t", false, "Show the total number of lines")
	statsCmd.Flags().BoolP("unique", "u", false, "Show the number of unique lines")
	statsCmd.Flags().BoolP("duplicate", "d", false, "Show the number of duplicate lines")
	statsCmd.Flags().BoolP("empty", "e", false, "Show the number of empty lines")
	statsCmd.Flags().BoolP("size", "s", false, "Show size of the file")
	statsCmd.Flags().BoolP("bytes", "b", false, "Show number of bytes of the file")

	statsCmd.Flags().BoolP("all", "a", false, "Show all the stats available")
}
