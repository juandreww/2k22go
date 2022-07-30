package main

import (
	"fmt"
)

type Car struct {
	Name string
	Age int
	PhoneNo int
}

func main() {
	c := Car{
		Name: 		"Adam", 
		Age:		24, 
		PhoneNo:	249899100}
	fmt.Println(c)
}