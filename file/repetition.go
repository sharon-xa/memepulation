package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type action func(index int)

// finding duplicates, and with every duplicate we call the action function
func (f *File) findDuplicates(fn action) {
	for i := 0; i < len(f.LinesArr)-1; i++ {
		for j := i + 1; j < len(f.LinesArr); j++ {
			if f.LinesArr[i] == f.LinesArr[j] {
				fn(i)
			}
		}
	}
}

func (f *File) NumOfRepeatedLines() {
	numberOfRepeatedLines := 0

	// increment the counter with every duplicate
	f.findDuplicates(func(i int) {
		numberOfRepeatedLines++
	})

	if numberOfRepeatedLines == 0 {
		fmt.Println("No repeated lines in the file.")
	}

	if numberOfRepeatedLines == 1 {
		fmt.Println("There's only one repeated line in the file.")
	}

	if numberOfRepeatedLines > 1 {
		fmt.Printf("There's %d repeated lines in this file.\n", numberOfRepeatedLines)
	}
}

func (f *File) ShowRepeatedLines() {
	// print the duplicated lines
	// TODO: you need to print something when there's no output (there is no duplicate)
	f.findDuplicates(func(index int) {
		fmt.Println(f.LinesArr[index])
	})
}

// TODO: if there's no duplicates you shouldn't run this method
func (f *File) NewFileWithNoDuplicates() {
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
	err := os.WriteFile(getFullFilePath(f.fileName), []byte(newContent), 0644)
	if err != nil {
		log.Println("couldn't create a file. ", err)
		return
	}
	fmt.Println("File Created Successfully.")
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
