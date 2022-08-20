package main

import (
	"fmt"
)

var horseNumbers = []int{889, 112, 334, 556, 838, 909, 990, 883, 723, 233}

func main() {
	fmt.Println("Ukuran maximum slice adalah:", cap(horseNumbers))
	fmt.Println("Kuda yang akan bertarung berjumlah:",len(horseNumbers))
	fmt.Println(horseNumbers)

	for k, _ := range horseNumbers {
		fmt.Println("Aku adalah kuda urutan ke", k)
	}

	for k, v := range horseNumbers {
		fmt.Printf("Kuda bernomor %d akan berada di urutan ke %d\n", v, k)
	}

}