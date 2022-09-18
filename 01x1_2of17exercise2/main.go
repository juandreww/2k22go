package main

import (
	"fmt"
)

var farm map[string]int

/*
	1. inisialisasi MAP, menggunakan key string value int
	2. print out map nya
	3. print key
	4. print key & value
	---
	1. peternakan & jumlah binatangnya
	2. ada ayam, sapi, ikan, kambing, anjing
*/

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