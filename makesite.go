package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
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

func mdToHtml(fileName string) {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	outputfileHtml := strings.Split(fileName, ".")[0] + ".html"
	md := markdown.ToHTML(fileContents, nil, nil)
	newFile, err := os.Create(outputfileHtml)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	_, err2 := newFile.WriteString("<!doctype html><html lang='en'><head><meta charset='utf-8'><title>Untitled Custom SSG</title></head><body>")
	if err2 != nil {
		panic(err)
	}
	defer newFile.Close()
	_, err3 := newFile.Write(md)
	if err3 != nil {
		panic(err)
	}
	defer newFile.Close()
	_, err4 := newFile.WriteString("</body></html>")
	if err4 != nil {
		panic(err)
	}
	newFile.Close()
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
				translatedFileData := fileData[count]
				newPara := para{Data: translatedFileData}
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
	//decare dir flag
	dirFlag := flag.String("dir", ".", "Directory to parse for txt files")

	//declare single file flag
	fileToRead := flag.String("file", "first-post.txt", "Single file to parse to html")

	//declare md flag
	mdFlag := flag.String("md", "testreadme.md", "Single md doc to be rendered to HTML")

	flag.Parse()

	//save mdToHTML result to variable
	if *mdFlag != "testreadme.md" {
		mdToHtml(*mdFlag)
		os.Exit(0)
	}
	//output all txt files in current directory
	directory := *dirFlag
	if *dirFlag != "." {
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			savePageToHtml(file.Name())
		}
		os.Exit(0)
	}
	//generate output file name
	if *fileToRead != "first-post.txt" {
		outputFile := *fileToRead
		savePageToHtml(outputFile)
		os.Exit(0)
	}

	//default operation if no flags passed
	fmt.Println("No flags passed, performing all operations on default values")
	mdToHtml(*mdFlag)
	files, err := ioutil.ReadDir(directory)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			savePageToHtml(file.Name())
		}
	outputFile := *fileToRead
	savePageToHtml(outputFile)
}
