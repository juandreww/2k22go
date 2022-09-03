package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(minilemon))
	mux.HandleFunc("/dog/", dog)
	mux.HandleFunc("/snoopy1", ServeSnoopy)
	fmt.Println("hello")

	http.ListenAndServe(":8080", mux)
}

func minilemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Mini Lemon toy found here")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func ServeSnoopy(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "snoopy1.jpg")
}