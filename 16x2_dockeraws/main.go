package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/", index)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oh yeah, I'm running on AWS.")
}
