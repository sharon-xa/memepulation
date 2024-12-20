package file

import (
	"log"
	"os"
	"strings"
)

type File struct {
	FileName    string
	fileSize    int64
	fileContent []byte
	linesArr    []string

	duplicates    int
	unique        int
	empty         int
	lineFrequency map[string]int
}

func ReadFile(path string) *File {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("can't read file: ", path, "\n", err)
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal("can't read file: ", path, "\n", err)
	}

	fileContent := strings.ToLower(string(file))

	arrOfLines := strings.Split(string(fileContent), "\n")
	lastArrIndex := arrOfLines[len(arrOfLines)-1]

	if lastArrIndex == "" || lastArrIndex == " " {
		arrOfLines = arrOfLines[:len(arrOfLines)-1]
	}

	pathArr := strings.Split(path, "/")
	fileName := pathArr[len(pathArr)-1]
	f := &File{
		FileName:    fileName,
		fileSize:    fileInfo.Size(),
		fileContent: []byte(fileContent),
		linesArr:    arrOfLines,
	}

	dup, uni, mt, lineFreq := f.analyzeLines()
	f.duplicates = dup
	f.unique = uni
	f.empty = mt
	f.lineFrequency = lineFreq

	return f
}

// analyzeLines analyzes the file lines and returns the count of duplicate, unique, and empty lines,
// along with the indexes of repeated lines.
func (f *File) analyzeLines() (duplicates int, unique int, empty int, lineFrequency map[string]int) {
	lineFrequency = make(map[string]int)

	for _, line := range f.linesArr {
		line = strings.TrimSpace(line)
		if line == "" {
			empty++
			continue
		}
		if lineFrequency[line] == 0 { // First occurrence of the line
			lineFrequency[line] = 1
			unique++
		} else if lineFrequency[line] > 0 { // Duplicate found
			lineFrequency[line] += 1
			duplicates++
		}
	}

	return duplicates, unique, empty, lineFrequency
}
