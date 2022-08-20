package main

import (
	"io"
	"net/http"
	"fmt"
	"math"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

type square struct {
	size int
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Round(math.Pi * c.radius * c.radius * 100) / 100
}

func (s square) areaof() {
	fmt.Println(s.size * s.size)
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{10}
	cc := circle{10}
	sq.areaof()
	printArea(cc)
	
	// http.HandleFunc("/", foo)
	// http.HandleFunc("/dog", bar)

	// fmt.Println("Serving...");
	// http.ListenAndServe(":8080", nil)
}