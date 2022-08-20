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
	x := []int{7, 14, 21}
	fmt.Println(x)
	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 0, 100)
	y = append(y,200)
	y = append(y,400)
	y = append(y,800)
	fmt.Println(y)
	for i, v := range y {
		fmt.Println(i, "-", v)
	}

	// http.HandleFunc("/", foo)
	// http.HandleFunc("/dog", bar)

	// fmt.Println("Serving...");
	// http.ListenAndServe(":8080", nil)
}