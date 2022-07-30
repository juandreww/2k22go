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

func main() {
	c := Car{
		Name: 		"Adam", 
		Age:		24, 
		PhoneNo:	249899100}
	c.Print()
}