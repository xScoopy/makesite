package main

import (
	"fmt"
	"io/ioutil"
)
type Header struct {
	Content string
}
type Paragraph struct {
	Section []string
}
// Read in the contents of the provided first-post.txt file
func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	return(string(fileContents))
}
//Edit the provided HTML template (`template.tmpl`)
//to display the contents of `first-post.txt`.
func updateTemplate(fileContents string) {

}
func main() {
	fmt.Println("Hello, world!")
	fmt.Print(readFile())
}
