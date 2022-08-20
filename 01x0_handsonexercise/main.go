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

func main() {
	aa := make(map[string]string)
	aa["James"] = "007";
	fmt.Println(aa["James"])


	p1 := person{"Nina", "Dobrev",}
	p2 := person{"Alicia", "Key",}
	fmt.Println(p1)
	fmt.Println(p2)

	a1 := struct{
		breed string
		name string
	}{
		"Holywings",
		"Shaun",
	}

	fmt.Println(a1)


	// http.HandleFunc("/", foo)
	// http.HandleFunc("/dog", bar)

	// fmt.Println("Serving...");
	// http.ListenAndServe(":8080", nil)
}