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

	1. buat struct person
	2. buat struct agent fields, dan embed person
	3. gunakan method pSpeak
	4. gunakan method secretAgentSpeak
	5. buat variabel type person
	6. buat variabel type secret agent
	7. panggil method Speak dari person
	8. print secret agent
	9. print speak dari variabel secret agent
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

func (fc oriflame) isoriflame() {
	if fc.isSkinCare == true {
		fmt.Println("\"", fc.fName, fc.lName, "\" is product of Oriflame")
	} else {
		fmt.Println("\"", fc.fName, fc.lName, "\" is not an Oriflame")
	}
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
	of.isoriflame()

	ofs := oriflame {
		facecream {
			"East Girl",
			"Normal Gel",
		},
		false,
	}
	ofs.isoriflame()
}