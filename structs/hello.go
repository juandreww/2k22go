package main

import (
	"fmt"
)

type Car interface {
	Drive()
	// FillGasoline() 
	Stop()
}

type Toyota struct {
	ToyotaModel string
}

type Honda struct {
	HondaModel string
}

func (h *Honda) Drive() {
	fmt.Println("Honda on the move")
	fmt.Println(h.HondaModel)
}

func (t *Toyota) Drive() {
	fmt.Println("Toyota on the go")
	fmt.Println(t.ToyotaModel)
}

func (t *Toyota) Stop() {
	fmt.Println("Toyota is stopping")
}

func newModel(arg string) Car {
	return &Toyota{arg}
}


// func main() {
// 	t := Toyota{"Yaris"}
// 	t.Drive()

// 	h := Honda{"Jazz"}
// 	h.Drive()
	
// 	nm := newModel("Harrier")
// 	nm.Drive()
// }