package main

import (
	"fmt"
)

var farm map[string]int

func main() {
	farm = map[string]int{}

	farm["dog"] = 1
	farm["chicken"] = 177
	farm["cow"] = 88
	farm["fish"] = 256
	farm["goat"] = 33

	fmt.Println("Jumlah anjing adalah", farm["dog"])
	fmt.Println()
	
	for k, _ := range farm {
		fmt.Println("Ada hewan", k)
	}

	fmt.Println()

	for k, v := range farm {
		fmt.Println("Jumlah hewan", k, "adalah", v)
	}
}