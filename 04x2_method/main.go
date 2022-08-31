package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url`"
)

type handlerlemon float64

func (p handlerlemon) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var p lemonhandler
	http.ListenAndServe(":8080", p)
}