package main

import (
	"fmt"
)

func main() {
	m1 := 3000000000000000000
	ltr := &m1
	fmt.Println(ltr, *ltr)
}