package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
