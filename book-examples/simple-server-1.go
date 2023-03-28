// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello, World!")
// 	})

// 	server := http.Server{
// 		Addr:    "127.0.0.1:5003",
// 		Handler: nil,
// 	}

// 	server.ListenAndServe()
// }