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
	person,
	allowedtobringgun bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says, "Good morning, James."`)
}

func main() {
	x := 7
	// started(x)
	fmt.Println(x)

	xi := []int{1,2,3,4}
	fmt.Println(xi)

	mp := map[string]int {
		"Boy" : 10,
		"Girl" : 10,
	}
	fmt.Println(mp)

	p1 := person{
		"Miss",
		"MoneyPenny",
	}
	fmt.Println(p1)
	p1.speak()
}