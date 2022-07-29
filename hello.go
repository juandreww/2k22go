package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := "mangtas"
	m2 := "mang"
	fmt.Println(strings.Contains(m1, m2))
}