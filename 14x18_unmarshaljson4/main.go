package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data string
	rcvd := `"Alice in Borderland"`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}
