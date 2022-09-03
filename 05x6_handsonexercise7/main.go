package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()
}