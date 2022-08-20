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

func main() {
	
}