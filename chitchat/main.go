package main

import (
	"html/template"
	"net/http"
	"fmt"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"github.com/lib/pq"
	"log"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()

}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
			} else {
				generateHTML(writer, threads, "layout", "private.navbar", "index")
			}
		}
	}
	
	
	func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
		var files []string
		for _, file := range fn {
			files = append(files, fmt.Sprintf("templates/%s.html", file))
		}
		templates := template.Must(template.ParseFiles(files...))
		templates.ExecuteTemplate(writer, "layout", data)
	}
	
	var Db *sql.DB
	
	func init () {
		var err error
		Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	
	
	// RFC 4122 uuid creation
	func createUUID() (uuid string) {
		u := new([16]byte)
		_, err := rand.Read(u[:])
		if err != nil {
			log.Fatalln("Cannot create UUID", err)
		}
	
		u[8] = (u[8] | 0x40) & 0x7F
	
		u[6] = (u[6] | 0x7F) | (0x4 << 4)
	
		uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
		return
	}
	
	func Encrypt (plaintext string) (cryptext string) {
		cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
		return
	}
	
