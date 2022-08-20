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

func (p woman) validateweight() string {
	if p.weight < 50.5 {
		return p.name + " is skinny"
	} else {
		return p.name + " is fat"
	}
}

func (p man) validateweight() string {
	if p.weight < 65.5 {
		return p.name + " is skinny"
	} else {
		return p.name + " is fat"
	}
}

type validation interface {
	validateheight() string
	validateweight() string
}

func myOpinion(v validation) {
	fmt.Println("Here is my thought...", v.validateheight())
	
}

func my2ndOpinion(v validation) {
	fmt.Println("Here is my 2nd thought...", v.validateweight())
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
		55,
	}
	fmt.Println(man.validateheight())
	fmt.Println(wman.validateheight())
	myOpinion(wman)
	myOpinion(man)
	fmt.Println(man.validateweight())
	fmt.Println(wman.validateweight())
	my2ndOpinion(wman)
	my2ndOpinion(man)
}