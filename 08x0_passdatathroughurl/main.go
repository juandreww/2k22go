package main

import (
	"fmt"
	"net/http"

)

func main() {
	http.HandleFunc("/", first)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func first(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("try")
	fmt.Fprintln(w, "try: " +v)
}