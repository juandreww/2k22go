package main

import (
	"fmt"
)

var x int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	allowedtobringgun bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (p secretAgent) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func main() {
	p1 := person{
		"Miss",
		"MoneyPenny",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"Donny",
			"Yen",
		},
		true,
	}
	fmt.Println(sa1)
	sa1.speak()
}