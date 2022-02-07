package main

import (
	"flag"
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

func savePage(newFileName string) {

}

func main() {
	dirFlag := flag.String("dir", "..", "Directory to parse for txt files")

	//use flag to get input data
	fileToRead := flag.String("file", "first-post.txt", "File to parse to html")
	flag.Parse()
	//output all txt files in current directory
	directory := *dirFlag
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			outputfileHtml := strings.Split(file.Name(), ".")[0] + ".html"
			txtFile := strings.Split(file.Name(), ".")[0]
			fileData := readFile(txtFile)
			header := fileData[0]
			//Setup paragraphs for body content
			var bodyContent []para
			for count := 1; count < len(fileData); count++ {
				newPara := para{Data: fileData[count]}
				bodyContent = append(bodyContent, newPara)
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
