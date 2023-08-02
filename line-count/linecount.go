// This program counts the number of lines in a specified file, of type html, css, or js.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run linecount.go <file_type> <file_path>")
		fmt.Println("Example: go run linecount.go css style.css")
		return
	}

	fileType := os.Args[1]
	filePath := os.Args[2]

	if fileType != "html" && fileType != "css" && fileType != "js" {
		log.Fatalf("Invalid file type: %s. Please specify 'html', 'css', or 'js'.", fileType)
		return
	}

	lineCount, err := countLinesInFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	fmt.Printf("Line count in %s file: %d\n", fileType, lineCount)
}

func countLinesInFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount, scanner.Err()
}
