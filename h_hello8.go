package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)

	// fmt.Println(c)
	// for i := range c {
	// 	fmt.Println(i)
	// }
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
	}()

	fmt.Println("Fin")
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Done")
}

