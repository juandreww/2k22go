package main

import (
	"fmt"
)

/*
	buat struct facecream
	fName North For Men
	lName Ultimate Balance

	buat struct oriflame
	facecream
	isSkinCare
*/

type facecream struct {
	fName string
	lName string
}

type oriflame struct {
	facecream
	isSkinCare bool
}

func (fc facecream) note() {
	fmt.Println("\"", fc.fName, fc.lName, "\" be your own beauty")
}

func main() {
	fs := facecream{
		"North for Men",
		"Ultimate Balance",
	}
	fmt.Println(fs)
	fs.note()
	
	of := oriflame {
		facecream {
			"Southern Woman",
			"Young Gel",
		},
		true,
	}

	of.note()
}