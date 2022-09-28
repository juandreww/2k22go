package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", prices.index)
	http.HandleFunc("/index/show", indexShow)
	http.HandleFunc("/index/create", indexCreateForm)
	http.HandleFunc("/index/create/process", indexCreateProcess)
	http.HandleFunc("/index/update", indexUpdateForm)
	http.HandleFunc("/index/update/process", indexUpdateProcess)
	http.HandleFunc("/index/delete/process", indexDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
