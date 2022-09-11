package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index2", index2)
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello from a docker container")
}

func index2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "aaa hello from a docker container")
}