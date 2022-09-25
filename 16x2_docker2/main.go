package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/index1", index)
	http.HandleFunc("/index2", index2)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from docker index 1")
}

func index2(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from docker index 2")
}
