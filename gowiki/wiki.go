package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func renderTemplate(w http.ResponseWriter, templ string, p *Page) {
	t, _ := template.ParseFiles(templ + ".html")
	t.Execute(w, p)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// ugly and cumbersome, templates/frameworks solve this
	// fmt.Fprintf(w, "<h1>Editing %s</h1>" + 
	// "form action=\"/save/%s\" method=\"POST\"" +
	// "textarea name=\"body\">%s</textarea><br>" +
	// "<input type=\"submit\" value=\"Save\">"+
	// "</form>",
	// p.Title, p.Title, p.Body)
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", viewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}