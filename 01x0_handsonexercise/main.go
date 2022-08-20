package main

import (
	"io"
	"net/http"
	"fmt"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", bar)

	fmt.Println("Serving...");
	http.ListenAndServe(":8080", nil)
}