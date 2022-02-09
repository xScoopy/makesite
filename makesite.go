package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)
type Content struct {
	Header string
	Paragraphs []para
}
type para struct {
	Data string
}
// Read in the contents of the provided first-post.txt file
func readFile(fileName string) []string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return (strings.Split(string(fileContents), "\n"))
}

func savePageToHtml(newFileName string) {
	if filepath.Ext(newFileName) == ".txt" {
		outputfileHtml := strings.Split(newFileName, ".")[0] + ".html"
		txtFile := newFileName
		fileData := readFile(txtFile)
		header := fileData[0]
		//Setup paragraphs for body content
		var bodyContent []para
		for count := 1; count < len(fileData); count++ {
			if fileData[count] != "" {
				fmt.Println(fileData[count])
				newPara := para{Data: fileData[count]}
				bodyContent = append(bodyContent, newPara)
			}
		}
		//initialize content struct for passing to template
		structuredContent := Content{Header: header, Paragraphs: bodyContent}
		templateParse := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
		newFile, err := os.Create(outputfileHtml)
		if err != nil {
			  panic(err)
		}
		templateParse.Execute(newFile, structuredContent)
	}
}

func main() {
	dirFlag := flag.String("dir", "..", "Directory to parse for txt files")

	//use flag to get input data
	fileToRead := flag.String("file", "first-post.txt", "Single file to parse to html")
	flag.Parse()
	//output all txt files in current directory
	directory := *dirFlag
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		savePageToHtml(file.Name())
	}
	//generate output file name
	outputFile := *fileToRead
	savePageToHtml(outputFile)

}
