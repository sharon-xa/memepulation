package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type action func(index int)

// this function find duplicates, and with every duplicate we call the action function
func (f *File) findDuplicates(fn action) {
	linesFrequency := make(map[string]int)

	for i, line := range f.LinesArr {
		line = strings.TrimSpace(line) // Trim spaces around the line
		if line == "" {
			continue // Skip empty lines
		}
		if linesFrequency[line] == 0 {
			// First occurrence of the line
			linesFrequency[line] = 1
		} else if linesFrequency[line] > 0 {
			// Duplicate found
			linesFrequency[line] += 1
			fn(i)
		}
	}
}

func (f *File) NumOfRepeatedLines() int {
	numberOfRepeatedLines := 0

	f.findDuplicates(func(i int) {
		numberOfRepeatedLines++
	})

	switch numberOfRepeatedLines {
	case 0:
		fmt.Println("No repeated lines in the file.")
	case 1:
		fmt.Println("There's only one repeated line in the file.")
	default:
		fmt.Printf("There's %d repeated lines in this file.\n", numberOfRepeatedLines)
	}

	return numberOfRepeatedLines
}

func (f *File) ShowRepeatedLines() {
	lineDuplicates := make(map[string]int)
	defer clear(lineDuplicates)

	f.findDuplicates(func(index int) {
		lineDuplicates[f.LinesArr[index]]++
	})

	fmt.Println("")
	if len(lineDuplicates) == 0 {
		fmt.Println("There are no duplicates")
		return
	}

	fmt.Println("")
	fmt.Println("If there's 2 names in a file the count will be 1 if 3 the count will be 2.")
	for line, count := range lineDuplicates {
		fmt.Printf("%s: %d\n", line, count)
	}

	fmt.Println("")
}

func (f *File) NewFileWithNoDuplicates() {
	if f.NumOfRepeatedLines() == 0 {
		fmt.Println("There are no duplicates in this file.")
		return
	}

	var newFileContent []string

	linesFrequency := make(map[string]int)

	for _, line := range f.LinesArr {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if linesFrequency[line] == 0 {
			linesFrequency[line] = 1
		}
	}

	for k := range linesFrequency {
		newFileContent = append(newFileContent, k)
	}

	newContent := strings.Join(newFileContent, "\n")

	// 0644: rw-r--r--
	fullPath := getFullFilePath(f.FileName)
	err := os.WriteFile(fullPath, []byte(newContent), 0644)
	if err != nil {
		log.Println("couldn't create a file.", err)
		return
	}
	fmt.Println("File Created Successfully.")
	fmt.Println("File path is:", fullPath)
}

func getFullFilePath(filename string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	newFileName := "new-" + filename
	fullPath := filepath.Join(homeDir, "Documents", newFileName)
	return fullPath
}
