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
	http.HandleFunc("/index", index)
	http.HandleFunc("/indextwo", indextwo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func indextwo(w http.ResponseWriter, r *http.Request) {
	str := "Here is another text from index two"
	fmt.Fprintln(w, str)
}
