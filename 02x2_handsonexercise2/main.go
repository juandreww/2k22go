package main

import (
	"fmt"
	// "log"
	// "os"
)

type Hotel struct {
	Name string
	Address string
	City string
	Zip string
	Region string
}

type Region struct {
	Name string
	Hotel []Hotel
}

type California struct {
	Region []Region
}

func main() {
	calf := Region {
		Name: "Southern",
	}
	fmt.Println(calf)
}