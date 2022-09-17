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
	mux.HandleFunc("/index", index)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}