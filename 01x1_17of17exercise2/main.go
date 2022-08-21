package main

import (
	"fmt"
)

func main() {
	s := "I'd rather eat the vegetables than meat"
	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	fmt.Println(string([]byte(s)[:14]))
	// fmt.Println(s)
	// fmt.Println(s)
	// fmt.Println(s)
}