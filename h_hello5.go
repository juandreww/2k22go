package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}

func superheavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super Heavy")
	}
}

func main() {
	go heavy()
	go superheavy()
	fmt.Println("Fin")
	time.Sleep(time.Second * 5)
}