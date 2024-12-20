package file

import "fmt"

func (f *File) GetTotalLinesNum() int {
	return len(f.linesArr)
}

func (f *File) GetUniqueLines() int {
	return f.unique
}

func (f *File) GetDuplicateLines() int {
	return f.duplicates
}

func (f *File) GetEmptyLines() int {
	return f.empty
}

func (f *File) GetFileSize() string {
	return fmt.Sprintf("%.2f KB", bytesToKilobytes(f.fileSize))
}

func (f *File) GetFileBytes() int64 {
	return f.fileSize
}

func (f *File) PrintAllStats() {
	fmt.Printf("\n📊 File statistics:\n")
	fmt.Printf("- Total lines: %d\n", len(f.linesArr))
	fmt.Printf("- Unique lines: %d\n", f.unique)
	fmt.Printf("- Duplicate lines: %d\n", f.duplicates)
	fmt.Printf("- Empty lines: %d\n", f.empty)
	fmt.Printf("- File size: %s\n", f.GetFileSize())
	fmt.Printf("- File bytes: %d\n\n", f.GetFileBytes())
}

func bytesToKilobytes(bytes int64) float64 {
	return float64(bytes) / 1024
}
