package main

import (
	"fmt"
)

func main() {
	m1 := "mang"
	m2 := "mang tas"
	fmt.Println(strings.Split(m2, " "), m1+m2)
}