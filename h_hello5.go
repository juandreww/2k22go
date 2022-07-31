package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Hello")
	}
}

func main() {
	go heavy()
	fmt.Println("Fin")
}