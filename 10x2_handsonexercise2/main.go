package main

import (
    "net/http"
    "io"
)

func main() {
    mux := http.DefaultServeMux
    mux.HandleFunc("/index", index)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        panic(err)
    }
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oh yeah, I'm running on AWS.")
}

