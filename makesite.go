package main

import (
	"fmt"
	"io/ioutil"
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
func readFile() []string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	return (strings.Split(string(fileContents), "\n"))
}


func main() {
	fmt.Println("Hello, world!")
	fileData := readFile()
	//Setup header
	header := fileData[0]
	//Setup para
	var bodyContent []para
	for count := 1; count < len(fileData); count++ {
		newPara := para{Data: fileData[count]}
		bodyContent = append(bodyContent, newPara)
	}
	structuredContent := Content{Header: header, Paragraphs: bodyContent}
	fmt.Print(structuredContent)
}
