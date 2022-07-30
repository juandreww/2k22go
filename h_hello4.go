package main

import "fmt"

func main() {
	fmt.Println("Hello 4")
	fmt.Println("Sample A: Continue")
	for i := 1; i < 20; i++ {
		if i % 3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("Sample B: Break")
	for i := 1; i < 20; i++ {
		if i % 3 == 0 {
			fmt.Printf("Break at seq: %v\n", i)
			break
		}
		fmt.Println(i)
	}
}