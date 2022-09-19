package main

import (
	"encoding/json"
	"fmt"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	mod := model{}
	fmt.Println(mod)

	bs, err := json.Marshal(mod)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(bs)
	fmt.Println("--a--")

	fmt.Println(string(bs))
}
