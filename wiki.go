package main

import (
	"fmt"
	"io/ioutil"
)

/*
The Body element is a []byte rather than string because that is the
type expected by the io libraries.
*/
type Page struct {
	Title string
	Body  []byte
}

/*
This method's signature reads: "This is a method named save that takes as its receiver p,
 a pointer to Page . It takes no parameters, and returns a value of type error."
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error ){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main () {
	myPage := &Page{Title: "Some title", Body: []byte("Some body goes here.")}
	myPage.save()
	printMyPage, _ := loadPage("Some title")
	fmt.Println(string(printMyPage.Body))
}
