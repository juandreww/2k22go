package main

import (
	"fmt"
)

type Person struct {
	name string
	location string
	fruits []string
}

func main() {
	p1 := Person{
		"Larasati",
		"Farm",
		[]string{"apple", "banana", "cherry", "grape", "pears",},
	}

	fmt.Println(p1)
	for k, v := range p1.fruits {
		fmt.Println("index ke", k, "isinya", v)
	}
}