package main

import (
	"fmt"
)

type CarFree struct {
	Model string
}

func main() {
	c := make(chan *CarFree, 3)

	go func() {
		c <- &CarFree{"1"}
		c <- &CarFree{"2"}
		c <- &CarFree{"3"}
		c <- &CarFree{"4"}
		close(c)
	}()

	// fmt.Println("Fin")
	for i := range c {
		fmt.Println(i.Model)
	}
	// fmt.Println("Done")
}

