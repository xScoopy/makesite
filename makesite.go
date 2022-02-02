package main

import (
	"flag"
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

func savePage(newFileName string) {

}

func main() {
	//use flag to get input data
	fileToRead := flag.String("file", "first-post.txt", "File to parse to html")
	flag.Parse()
    
	//generate output file name
	outputFileMinusExt := strings.Split(*fileToRead, ".")[0]
	outputFile := outputFileMinusExt + ".html"

	fileData := readFile(*fileToRead)
	//Setup header
	header := fileData[0]
	//Setup paragraphs for body content
	var bodyContent []para
	for count := 1; count < len(fileData); count++ {
		newPara := para{Data: fileData[count]}
		bodyContent = append(bodyContent, newPara)
	}

	//initialize content struct for passing to template
	structuredContent := Content{Header: header, Paragraphs: bodyContent}
	
	//parse template and write to a new html file with data injected to template
	templateParse := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFile, err := os.Create(outputFile)
	if err != nil {
		  panic(err)
	}
	templateParse.Execute(newFile, structuredContent)
}
