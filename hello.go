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
	c := Car{"Adam", 24, 9899100}
	fmt.Println(c)
}