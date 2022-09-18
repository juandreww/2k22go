package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("abc")
	fs := http.FileServer(http.Dir("."))
	log.Fatal(http.ListenAndServe(":8080", fs))
}