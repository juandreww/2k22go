package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Fprintln(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}
