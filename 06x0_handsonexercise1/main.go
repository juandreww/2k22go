package main

import (
	"fmt"
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(minilemon))
	mux.Handle("/dog/", http.HandlerFunc(dog))
	mux.HandleFunc("/snoopy1", ServeSnoopy)
	fmt.Println("hello")

	http.ListenAndServe(":8080", mux)
}

func minilemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Mini Lemon toy found here")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		panic(err)
	}
	data := "here is the dog"

	tpl.ExecuteTemplate(w, "dog.gohtml", data)
}

func ServeSnoopy(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "snoopy1.jpg")
}