package main

import (
	"fmt"
)

type gator int
var x int

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	x = 42
	fmt.Println()
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// x = g1 + 1
	// 13 because the type is gator not int
	x = int(g1) + x
	fmt.Println("new value of x")
	fmt.Println(x)
}