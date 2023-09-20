package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	server := http.Server{
		Addr:    "127.0.0.1:8999",
		Handler: mux,
	}
	server.ListenAndServe()
}
