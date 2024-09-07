package file

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type File struct {
	FileName       string
	FileContent    []byte
	LinesArr       []string
	LinesArrLength int
}

func ReadFile(path string) *File {
	file, err := os.ReadFile(path)
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
	return &File{
		FileName:       fileName,
		FileContent:    []byte(fileContent),
		LinesArr:       arrOfLines,
		LinesArrLength: len(arrOfLines),
	}
}

func (f *File) GetFileContent() string {
	return string(f.FileContent)
}

func (f *File) PrintFileContent() {
	fmt.Print(string(f.FileContent))
}
