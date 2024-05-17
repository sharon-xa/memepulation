package file

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type File struct {
	fileName    string
	fileContent []byte
	LinesArr    []string
}

func ReadFile(path string) *File {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("can't read file: ", path, "\n", err)
	}

	arrOfLines := strings.Split(string(fileContent), "\n")
	if arrOfLines[len(arrOfLines)-1] == "" || arrOfLines[len(arrOfLines)-1] == " " {
		arrOfLines = arrOfLines[:len(arrOfLines)-1]
	}
	pathArr := strings.Split(path, "/")
	fileName := pathArr[len(pathArr)-1]
	return &File{
		fileName:    fileName,
		fileContent: fileContent,
		LinesArr:    arrOfLines,
	}
}

func (f *File) GetFileContent() string {
	return string(f.fileContent)
}

func (f *File) PrintFileContent() {
	fmt.Print(string(f.fileContent))
}
