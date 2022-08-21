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
}

type sedan struct {

}

func main() {
	fmt.Println("hello")
}