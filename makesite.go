package main

import (
	"html/template"
	"io/ioutil"
	"os"
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


func main() {
	fileData := readFile("first-post.txt")
	//Setup header
	header := fileData[0]
	//Setup para
	var bodyContent []para
	for count := 1; count < len(fileData); count++ {
		newPara := para{Data: fileData[count]}
		bodyContent = append(bodyContent, newPara)
	}
	structuredContent := Content{Header: header, Paragraphs: bodyContent}
	

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFile, err := os.Create("first-post.html")
	if err != nil {
		  panic(err)
	}
	t.Execute(newFile, structuredContent)
}
