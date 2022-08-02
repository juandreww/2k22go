package main

import (
	"fmt"
)

func Select(c chan int, quits chan struct{}) {
	for {
		select {
		case <- c:
			fmt.Println("received")
		case <- quits:
			fmt.Println("Quit")
			break
		}
	}
}

func main() {
	c := make(chan int)
	quits := make(chan struct{})
	go Select(c, quits)

	go func() {
		c <- 1
	}()
}