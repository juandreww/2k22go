package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hallo")
	}()
	fmt.Println("Fin")
}

