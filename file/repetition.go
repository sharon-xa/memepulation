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
	for i := 0; i < len(f.LinesArr); i++ {
		for j := i + 1; j < len(f.LinesArr); j++ {
			if f.LinesArr[i] == f.LinesArr[j] {
				fn(i)
				break
			}
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

	for _, line := range f.LinesArr {
		// Check if the name exist in the new array
		exist := false
		for _, str := range newFileContent {
			if line == str {
				exist = true
			}
		}
		if !exist {
			newFileContent = append(newFileContent, line)
		}
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
