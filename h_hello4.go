package main

import "fmt"

func main() {
	fmt.Println("Hello 4")
	fmt.Println("Sample A")
	for i := 1; i < 20; i++ {
		if i % 3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("Sample B")
}