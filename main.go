package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", myMiddleWare(myHandler))

	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello from handler")
	w.WriteHeader(http.StatusOK)
}

func myMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello from middleware")
		next(w, r )
	}
}