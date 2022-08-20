package main

import (
	"fmt"
)

/*
	1. buatlah sebuah interface yang type di implementasi oleh person dan secretAgent
	2. deklarasi sebuah fungsi dengan parameter interface tsb
	3. panggil function tersebut di main dan value nya adalah type person
	4. panggil function tersebut di main dan value nya adalah secretAgent
*/

type man struct {
	name string
	height int
	weight float64
}

type woman struct {
	name string
	height int
	weight float64
}

func (p man) validateheight() string {
	if p.height < 170 {
		return p.name + " is short"
	} else {
		return p.name + " is tall"
	}
}

func (p woman) validateheight() string {
	if p.height < 155 {
		return p.name + " is short"
	} else {
		return p.name + " is tall"
	}
}

type validation interface {
	validateheight() string
}

func myOpinion(v validation) {
	fmt.Println("Here is my thought...", v.validateheight())
}

func main() {
	man := man{
		"Javier",
		180,
		64.5,
	}

	wman := woman{
		"Helena",
		150,
		45,
	}
	fmt.Println(man.validateheight())
	fmt.Println(wman.validateheight())
	myOpinion(wman)
	myOpinion(man)

}