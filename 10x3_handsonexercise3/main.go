package main

import (
    "net/http"
    // "io"
)

func main() {
    mux := http.DefaultServeMux
    mux.HandleFunc("/index", index)
    err := http.ListenAndServe(":8080", nil)

    check(err)
    
}

func check(err error) {
    if err != nil {
        panic(err)
    } 
}

