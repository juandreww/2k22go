package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/index", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is index: ", r.Method)

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}