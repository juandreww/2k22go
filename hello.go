package main

import (
	"fmt"
)

type Car interface {
	Drive()
	FillGasoline() 
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


func main() {

	
}