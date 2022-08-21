package main

import (
	"fmt"
)

type gator int
var x int

func (g1 gator) greeting() {
	fmt.Println("I am gator", g1)
}

func main() {
	var g1 gator
	g1 = 42
	g1.greeting()
}