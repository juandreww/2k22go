package main

import (
	"fmt"
) 

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()
	
	val := <- c

	fmt.Println(&c, c, val);
}

