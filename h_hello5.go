package main

import (
	"fmt"
	"time"
)

func heavy() {
	time.Sleep(time.Second * 5)
}

func main() {
	go heavy()
	fmt.Println("Fin")
}