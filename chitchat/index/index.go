package index

import (
	"html/template"
	"net/http"
	"fmt"
)
func index(w http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(w, request)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
			} else {
				generateHTML(w, threads, "layout", "private.navbar", "index")
			}
		}
	}
	
	
	func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
		var files []string
		for _, file := range fn {
			files = append(files, fmt.Sprintf("templates/%s.html", file))
		}
		templates := template.Must(template.ParseFiles(files...))
		templates.ExecuteTemplate(w, "layout", data)
	}
