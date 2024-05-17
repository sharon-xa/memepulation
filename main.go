package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sharon-xa/text-manipulation/file"
)

var pl = fmt.Println

func usage() {
	pl("")
	pl("This program helps you preform some operations on text files.")
	pl("Usage:")
	pl("    text-manipulation [file]")
	pl("")
	pl("Available Operations:")
	pl("    s   show file content.")
	pl("    r   show number of repeated lines.")
	pl("    sr  show repeated lines.")
	pl("    n  	create new file with no duplicates.")
	pl("    h   see help.")
	pl("    q   exit out of the program.")
}

func main() {
	usage()
	filename := readArgument()
	file := file.ReadFile(filename)

infinitLoop:
	for {
		input := bufio.NewScanner(os.Stdout)
		fmt.Print("\nEnter operation: ")
		input.Scan()

		switch input.Text() {
		case "s":
			file.PrintFileContent()

		case "r":
			file.NumOfRepeatedLines()

		case "sr":
			file.ShowRepeatedLines()

		case "n":
			file.NewFileWithNoDuplicates()

		case "h":
			usage()

		case "q":
			pl("exiting...")
			break infinitLoop

		default:
			pl("Unknown operation! try again.")
		}
	}
	pl("GoodBye")
}
