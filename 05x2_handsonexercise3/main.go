package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(home))
	mux.Handle("/warehouse", http.HandlerFunc(warehouse))
	mux.Handle("/factory", http.HandlerFunc(factory))
	mux.Handle("/shop", http.HandlerFunc(shop))


	http.ListenAndServe(":8080", mux)
}