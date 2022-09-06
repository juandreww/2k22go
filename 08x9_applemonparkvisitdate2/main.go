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
	mux.HandleFunc("/submitarrival", submitarrival)
	mux.HandleFunc("/bar", bar)
	mux.HandleFunc("/thankyou", thankyou)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form submission here
	http.Redirect(w, req, "/thankyou", http.StatusSeeOther)
}

func thankyou(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at thankyou:", req.Method)
	tpl.ExecuteTemplate(w, "thanyou.gohtml", nil)
}

func submitarrival(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at submitarrival:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}