package main

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/submitarrival", submitarrival)
	http.ListenAndServe(":8080", nil)
}

func submitarrival(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln("Hi, Thank you for visiting")
}