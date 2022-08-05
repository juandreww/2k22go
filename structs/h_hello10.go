package main

// import (
// 	"fmt"
// 	"os"
// )

// func Select(c chan int, quits chan struct{}) {
// 	for {
// 		select {
// 		case <- c:
// 			fmt.Println("received")
// 		case <- quits:
// 			fmt.Println("Quit")
// 			os.Exit(0)
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quits := make(chan struct{})
// 	go Select(c, quits)

// 	go func() {
// 		c <- 1
// 		quits <- struct{}{}
// 	}()

// 	c <- 1
// 	select{}
// }