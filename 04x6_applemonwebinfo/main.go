package main

import (
	// "fmt"
	"net/http"
	"html/template"
	"net/url"
	"log"
)

type lemonhandler float64

func (p lemonhandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method		string
		URL			*url.URL
		Submissions map[string][]string
		Header		http.Header
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
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