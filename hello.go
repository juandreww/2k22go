package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := "mang"
	m2 := "mang tas"
	fmt.Println(strings.Split(m2, " "), m1+m2)
}