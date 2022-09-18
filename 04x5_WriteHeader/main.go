package main

import (
	"fmt"
	"net/http"
)

type lemonhandler float64

func (p lemonhandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Toon-Lemon", "this is Toon Lemon Shop")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

// var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseFiles("index.gohtml"))
// }
func main() {
	var p lemonhandler
	http.ListenAndServe(":8080", p)
}