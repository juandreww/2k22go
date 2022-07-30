package main

import "fmt"

func main() {
	fmt.Println("Hello UK")

	f := true
	flag := &f

	fmt.Println(*flag, flag, &flag)
	if *flag {
		fmt.Println("True")
	} else if !*flag {
		fmt.Println("True")
	} else {

	}
}