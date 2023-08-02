// First attempt at get docx word count file in Go, not working yet, overestimates number of words.

package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func docxmain() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run docxwordcount.go <file_path>")
		return
	}
	filePath := os.Args[1]
	wordCount, err := countWordsInDocxFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	fmt.Printf("Word count in file: %d\n", wordCount)
}

func countWordsInDocxFile(filePath string) (int, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	if strings.HasSuffix(filePath, ".docx") {
		// Open the docx file using zip.Reader
		reader, err := zip.NewReader(strings.NewReader(string(content)), int64(len(content)))
		if err != nil {
			return 0, err
		}
		var docXML string
		for _, file := range reader.File {
			if file.Name == "word/document.xml" {
				f, err := file.Open()
				if err != nil {
					return 0, err
				}
				defer f.Close()

				b, err := ioutil.ReadAll(f)
				if err != nil {
					return 0, err
				}

				docXML = string(b)
				break
			}
		}
		wordCount := countWordsInXML(docXML)
		return wordCount, nil
	}

	return 0, fmt.Errorf("unsupported file format")
}

func countWordsInXML(xmlContent string) int {
	const (
		startTag = "<w:t"
		endTag   = "</w:t>"
	)

	wordCount := 0
	for {
		startIndex := strings.Index(xmlContent, startTag)
		if startIndex == -1 {
			break
		}
		endIndex := strings.Index(xmlContent, endTag)
		if endIndex == -1 || endIndex < startIndex {
			break
		}
		wordContent := xmlContent[startIndex : endIndex+len(endTag)]
		wordContent = strings.TrimPrefix(wordContent, startTag)
		wordContent = strings.TrimSuffix(wordContent, endTag)
		wordContent = strings.TrimSpace(wordContent)
		words := strings.Fields(wordContent)
		wordCount += len(words)
		xmlContent = xmlContent[endIndex+len(endTag):]
	}
	return wordCount
}
