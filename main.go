package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sharon-xa/memepulation/file"
)

var pl = fmt.Println

func usage() {
	pl("\nThis program helps you perform operations on text files.")
	pl("Usage:")
	pl("    text-manipulation [flags] [file]\n")
	pl("Flags:")
	flag.PrintDefaults() // Automatically prints the flags you've defined
	pl("\nExamples:")
	pl("    text-manipulation -s file.txt       # Show file content")
	pl("    text-manipulation -sr file.txt      # Show repeated lines")
	pl("    text-manipulation -n file.txt       # Create a new file without duplicates")
	pl("    text-manipulation -stats file.txt   # Show statistics about the file content")
	pl("")
}

func main() {
	showContent := flag.Bool("s", false, "Show file content")
	numRepeated := flag.Bool("r", false, "Show number of repeated lines")
	showRepeated := flag.Bool("sr", false, "Show repeated lines")
	createNew := flag.Bool("n", false, "Create a new file with no duplicates")
	stats := flag.Bool("stats", false, "Show file content statistics")
	help := flag.Bool("h", false, "Display help")

	flag.Usage = usage
	flag.Parse()

	if *help || len(flag.Args()) == 0 {
		usage()
		return
	}

	filename := flag.Arg(0)
	file := file.ReadFile(filename)
	if file == nil {
		log.Fatalln("Failed to read the file.")
		return
	}

	switch {
	case *showContent:
		file.PrintFileContent()

	case *numRepeated:
		file.NumOfRepeatedLines()

	case *showRepeated:
		file.ShowRepeatedLines()

	case *createNew:
		file.NewFileWithNoDuplicates()

	case *stats:
		file.ShowFileStats()

	default:
		pl("Unknown operation! Use -h to see available options.")
	}
}
