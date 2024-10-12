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

	fmt.Println("\n-------------------------")
	switch numberOfRepeatedLines {
	case 0:
		fmt.Println("✅ No repeated lines found in the file.")
	case 1:
		fmt.Println("⚠️  There is 1 repeated line in the file.")
	default:
		fmt.Printf("⚠️  There are %d repeated lines in the file.\n", numberOfRepeatedLines)
	}
	fmt.Print("-------------------------\n\n")

	return numberOfRepeatedLines
}

func (f *File) ShowRepeatedLines() {
	lineDuplicates := make(map[string]int)
	defer clear(lineDuplicates)

	f.findDuplicates(func(index int) {
		lineDuplicates[f.LinesArr[index]]++
	})

	fmt.Println("\n-------------------------")
	if len(lineDuplicates) == 0 {
		fmt.Println("✅ No duplicate lines found.")
		fmt.Print("-------------------------\n\n")
		return
	}

	fmt.Print("🔁 Repeated lines and their counts:\n\n")
	for line, count := range lineDuplicates {
		fmt.Printf("- \"%s\": repeated %d time(s)\n", line, count)
	}
	fmt.Print("-------------------------\n\n")
}

func (f *File) NewFileWithNoDuplicates() {
	if f.NumOfRepeatedLines() == 0 {
		fmt.Print("✅ No duplicates found. No new file created.\n\n")
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
	fullPath := getFullFilePath(f.FileName)
	err := os.WriteFile(fullPath, []byte(newContent), 0644)
	if err != nil {
		log.Println("❌ Failed to create the new file:", err)
		return
	}

	fmt.Println("✅ New file created successfully.")
	fmt.Printf("📁 File path: %s\n\n", fullPath)
}

func (f *File) ShowFileStats() {
	totalLines := len(f.LinesArr)
	emptyLines := 0
	uniqueLines := 0
	duplicateLines := 0
	linesFrequency := make(map[string]int)

	for _, line := range f.LinesArr {
		line = strings.TrimSpace(line)
		if line == "" {
			emptyLines++
			continue
		}
		if linesFrequency[line] == 0 {
			linesFrequency[line] = 1
			uniqueLines++
		} else {
			duplicateLines++
		}
	}

	fmt.Printf("\n📊 File statistics:\n")
	fmt.Printf("- Total lines: %d\n", totalLines)
	fmt.Printf("- Empty lines: %d\n", emptyLines)
	fmt.Printf("- Unique lines: %d\n", uniqueLines)
	fmt.Printf("- Duplicate lines: %d\n\n", duplicateLines)
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
