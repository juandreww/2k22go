package main

import (
	"io"
	// "fmt"
	"net/http"

)

func main() {
	http.HandleFunc("/", first)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func first(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("try")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="try">
	 <input type="submit">
	</form>
	<br>`+v)
}