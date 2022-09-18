package main

import (
	"fmt"
)

type gator int
type flamingo bool
var x int

func (p gator) greeting() string {
	return "I am gator"
}

func (p flamingo) greeting() string {
	return "I am the flamingo"
}

type swampCreature interface {
	greeting() string
}

func bayou(p swampCreature) {
	fmt.Println("do you know this ..." + p.greeting())
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

	bayou(g1)
	bayou(f1)
}