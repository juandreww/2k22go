package main

import (
	"fmt"
)

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println("hello")
}