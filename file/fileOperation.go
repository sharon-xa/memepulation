package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (f *File) PrintFileContent() {
	fmt.Print(string(f.fileContent))
}

func (f *File) PrintRepeatedLinesNumber() {
	fmt.Println("\n-------------------------")
	switch f.duplicates {
	case 0:
		fmt.Println("✅ No repeated lines found in the file.")
	case 1:
		fmt.Println("⚠️  There is 1 repeated line in the file.")
	default:
		fmt.Printf("⚠️  There are %d repeated lines in the file.\n", f.duplicates)
	}
	fmt.Print("-------------------------\n\n")
}

// TODO: lineFrequency is not lineDuplicates
// if a line exists 2 times we should show the user that
// this line is repeated one time not two times
func (f *File) PrintRepeatedLines() {
	fmt.Println("\n-------------------------")
	if f.duplicates == 0 {
		fmt.Println("✅ No duplicate lines found.")
		fmt.Print("-------------------------\n\n")
		return
	}

	fmt.Print("🔁 Repeated lines and their counts:\n\n")
	for line, count := range f.lineFrequency {
		if count > 1 {
			fmt.Printf("- \"%s\": repeated %d time(s)\n", line, count)
		}
	}
	fmt.Print("-------------------------\n\n")
}

// TODO: finish the create or new subcommand
func (f *File) NewFileWithNoDuplicates() {
	if f.duplicates == 0 {
		fmt.Print("✅ No duplicates found. No new file created.\n\n")
		return
	}

	var newFileContent []string
	linesFrequency := make(map[string]int)

	for _, line := range f.linesArr {
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
	fullPath := getFullFilePath(f.FileName)
	err := os.WriteFile(fullPath, []byte(newContent), 0644)
	if err != nil {
		log.Println("❌ Failed to create the new file:", err)
		return
	}

	fmt.Println("✅ New file created successfully.")
	fmt.Printf("📁 File path: %s\n\n", fullPath)
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
