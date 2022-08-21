package main

import (
	"fmt"
)

type Person struct {
	fName string
	lName string
}

func (p Person) walk() string {
	return p.fName + "is walking"
}

func main() {
	p1 := Person {
		"Donny",
		"Yen",
	}
	fmt.Sprintln(p1.walk())
}