package main

import (
	"fmt"
)

type gator int
type flamingo bool
var x int

func (p gator) greeting() {
	fmt.Println("I am gator", p)
}

func (p flamingo) greeting() {
	fmt.Println("I am the flamingo", p)
}

func main() {
	var g1 gator
	g1 = 42
	g1.greeting()

	var f1 flamingo
	f1 = false
	fmt.Println("I am not beautiful", f1)

	f1 = true
	f1.greeting()
}