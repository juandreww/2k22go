package main

import (
	"fmt"
	"net/http"
)

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

}