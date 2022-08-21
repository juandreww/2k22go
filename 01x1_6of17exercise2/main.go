package main

import (
	"fmt"
)

type Person struct {
	fName string
	lName string
}

func (p Person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

func main() {
	p1 := Person {
		"Donny",
		"Yen",
	}
	s := p1.walk()
	fmt.Println(s)
}