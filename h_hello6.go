package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hallo")
	}()
	func() {
		fmt.Println("World")
	}()
	fmt.Println("Fin")
}

