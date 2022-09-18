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
	fmt.Println(string([]byte(s)[10:22]))
	fmt.Println(string([]byte(s)[14:]))

	for _, v := range []byte(s) {
		fmt.Println(string(v))
	}
	// fmt.Println(s)
	// fmt.Println(s)
	// fmt.Println(s)
}