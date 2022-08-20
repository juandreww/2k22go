package main

import (
	"fmt"
)

var x int

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
}