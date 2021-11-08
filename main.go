package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 660)
}

func loadpage(title string) {

}
func editHandle() {

}
func main() {
	fmt.Println("hello, world")
}
