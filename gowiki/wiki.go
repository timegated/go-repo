package main

import (
	"fmt"
	"os"
	"log"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	err := os.WriteFile(filename, p.Body, 0600)
	if err !=nil {
		log.Fatal(err)
	}
	return err
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main () {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}