package main

import (
	"fmt"
)

type Printer struct {
	brand string
	model string
}

func main() {
	p1 := Printer{
		"Epson",
		"EP-200",
	}
	fmt.Println(p1)
}