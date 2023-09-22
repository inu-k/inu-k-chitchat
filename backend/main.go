package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"
)

type Response struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	jsonData := Response{
		Name:    "TestName",
		Message: "TestMessage",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseJson, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	c := cors.Default()
	handler := c.Handler(mux)
	mux.HandleFunc("/index", index)

	server := http.Server{
		Addr:    "127.0.0.1:8999",
		Handler: handler,
	}
	server.ListenAndServe()
}