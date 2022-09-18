package main

import (
	"fmt"
	"net/http"
)

type lemon float64

func (p lemon) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey, anything you want to order?")
}

func main() {
	var p lemon
	http.ListenAndServe(":8080", p)
}