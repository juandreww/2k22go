package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type lemonhandler float64

func (p lemonhandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var p lemonhandler
	http.ListenAndServe(":8080", p)
}