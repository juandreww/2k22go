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
}

type sedan struct {
	vehicle
	luxury bool
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

func (p truck) transportationDevice() string {
	return "The truck is four wheel"
}

func (p sedan) transportationDevice() string {
	return "The sedan is luxurious"
}

type transportation interface {
	transportationDevice() string
}

func report(p transportation) {
	fmt.Println(p.transportationDevice())
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
	report(trk)

	sdn := sedan {
		vehicle {
			"two doors",
			"red",
		}, false,
	}

	fmt.Println(sdn.isluxury())
	fmt.Println(sdn)
	report(sdn)


}