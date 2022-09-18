package main

import (
	"fmt"
	"net/http"
	"io"
)

type handler1 int

func (p handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>My Name is Andrew</h1>")
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>here is the index</h1>")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>here is the dog</h1>")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>my name is Andrew</h1>")
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/dog", dog)
	mux.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", mux)
}