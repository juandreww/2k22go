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
}