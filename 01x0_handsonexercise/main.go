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

type person struct {
	fName string
	lName string
}

func (p person) speak() {
	fmt.Println("Speak of", p.fName)
}

func main() {
	p1 := person{"Nina", "Dobrev",}
	p1.speak()

	// http.HandleFunc("/", foo)
	// http.HandleFunc("/dog", bar)

	// fmt.Println("Serving...");
	// http.ListenAndServe(":8080", nil)
}

func meaning() int {
	return 42
}