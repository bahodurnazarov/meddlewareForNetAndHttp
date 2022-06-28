package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandler)

	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hi")
	w.WriteHeader(http.StatusOK)
}