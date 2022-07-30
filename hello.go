package main

import (
	"fmt"
)

type Car struct {
	Name string
	Age int
	PhoneNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Announce() {
	fmt.Println("car is discounted")
}

func (c Car) getPhone() int {
	return c.PhoneNo
}

func main() {
	c := Car{
		Name: 		"Adam", 
		Age:		24, 
		PhoneNo:	249899100}
	c.Print()
	c.Announce()
	d := c.getPhone()
	fmt.Println(d)
}