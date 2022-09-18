package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type handlerlemon float64

func (p handlerlemon) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		URL *url.URL
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var p handlerlemon
	http.ListenAndServe(":8080", p)
}