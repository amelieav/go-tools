// Simple program that counts the number of words in a .txt file.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run txtwordcount.go <file_path>")
		return
	}

	filePath := os.Args[1]
	wordCount, err := countWordsInFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	fmt.Printf("Word count in file: %d\n", wordCount)
}

func countWordsInFile(filePath string) (int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	words := strings.Fields(string(content))
	return len(words), nil
}
