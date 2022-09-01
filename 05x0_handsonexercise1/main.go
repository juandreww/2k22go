package main

import (
	"fmt"
	"net/http"
)

type handler1 int

func (p handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>My Name is Andrew</h1>")
}

func main() {
	fmt.Println("main")
}