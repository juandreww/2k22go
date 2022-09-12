package main

import (
    "net/http"
    "io"
)

func main() {
    mux := http.DefaultServeMux
    mux.HandleFunc("/index", index)
    mux.HandleFunc("/index2", index2)
    mux.HandleFunc("/", index)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        panic(err)
    }
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oh yeah, I'm running on AWS.")
}

func index2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oh yeah, I'm running on AWS.")
}

