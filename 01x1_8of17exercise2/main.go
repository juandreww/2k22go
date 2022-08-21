package main

import (
	"fmt"
)

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
	transportationDevice()
}

type sedan struct {
	vehicle
	luxury bool
	transportationDevice()
}

func (p truck) isFourWheel() string {
	return fmt.Sprintln("The truck is four wheel:", p.fourWheel)
}

func (p sedan) isluxury() string {
	if p.luxury == true {
		return "sedan with " + p.doors + " and " + p.color + " color is luxury"
	} else {
		return "sedan with " + p.doors + " and " + p.color + " color is not luxury"
	}
}

type transportation interface {
	transportationDevice() string
}

func (p transportation) report() string {
	
}

func main() {
	trk := truck {
		vehicle {
			"four doors",
			"blue",
		}, true,
	}

	fwd := trk.isFourWheel()
	fmt.Println(fwd)
	fmt.Println(trk)

	sdn := sedan {
		vehicle {
			"two doors",
			"red",
		}, false,
	}

	fmt.Println(sdn.isluxury())
	fmt.Println(sdn)
}