package main

import (
	"fmt"
	"net/http"
	"html/template"
	"io"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(minilemon))
	mux.Handle("/dog/", http.HandlerFunc(dog))
	fmt.Println("hello")

	http.ListenAndServe(":8080", mux)
}

func minilemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Mini Lemon toy found here")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Here is the index</h1>")
	data := "here is the dog"

	tpl.ExecuteTemplate(w, "dog.gohtml", data)
	http.ServeFile(w, r, "snoopy1.jpg")
}