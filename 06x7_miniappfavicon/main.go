package main

import (
	// "html/template"
	"log"
	"net/http"
	"fmt"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	// mux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)

	fmt.Fprintln(w, "go look at your terminal")
	// HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}