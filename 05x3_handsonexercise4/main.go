package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(home))
	mux.Handle("/warehouse", http.HandlerFunc(warehouse))
	mux.Handle("/factory", http.HandlerFunc(factory))
	mux.Handle("/shop", http.HandlerFunc(shop))


	http.ListenAndServe(":8080", mux)
}

func warehouse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, youre in our warehouse.")
}

func factory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, youre in our factory.")
}

func shop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, youre in our shop.")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, youre in our home.")
}