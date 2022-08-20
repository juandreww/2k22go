package main

import (
	"fmt"
)

type Person struct {
	fName string
	lName string
}

func main() {
	p1 := Person{
		"Andika",
		"Wijaya",
	}
	fmt.Println(p1)
}