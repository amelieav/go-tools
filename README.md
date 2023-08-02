# go-tools
### Simple tools coded in Golang.

linecount.go counts the number of lines in a given css, js or html file. 
Run with: ```go run linecount.go <file_type> <file_path>```
For example: ```go run linecount.go css style.css```

#### To-Do: Be able to pass in a folder and return analysis of files within it.

txtwordcount counts the number of words in a .txt extension file.
Run with: ```go run txtwordcount.go``` in command line within txt-word-count folder.

docxwordcount is first attempt at counting number of words in .docx extension file. UNFINISHED. It does not work at this time.
