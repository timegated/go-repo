// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello, JEFF")
// 	})

// 	server := http.Server{
// 		Addr:              "127.0.0.1:5003",
// 		Handler:           nil,
// 		ReadTimeout:       10 * time.Second,
// 		WriteTimeout:      10 * time.Second,
// 		IdleTimeout:       60 * time.Second,
// 		MaxHeaderBytes:    1 << 20,
// 		ErrorLog:          nil,
// 		BaseContext:       nil,
// 		ConnContext:       nil,
// 		TLSConfig:         nil,
// 		TLSNextProto:      nil,
// 		ConnState:         nil,
// 	}

// 	server.ListenAndServe()
// }