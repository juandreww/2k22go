package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	mod := model{
		State: true,
		Pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}
	fmt.Println(mod)

	bs, err := json.Marshal(mod)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(bs)
	fmt.Println("--a--")

	fmt.Println(string(bs))
	fmt.Println("--b--")

	os.Stdout.Write(bs)
}
