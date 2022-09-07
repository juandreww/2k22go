package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/submit", submit)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	
}