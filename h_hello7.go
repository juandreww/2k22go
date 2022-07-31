package main

// import (
// 	"fmt"
// 	"time"
// ) 

// func main() {
// 	c := make(chan int)

// 	go func() {
// 		c <- 1
// 	}()
	
// 	val := <- c
// 	fmt.Println(&c, c, val)

// 	go func() {
// 		c <- 3
// 	}()
	
// 	time.Sleep(time.Second * 2)
// 	val = <- c
// 	fmt.Println(&c, c, val)
// 	fmt.Println("Fin")
// }

