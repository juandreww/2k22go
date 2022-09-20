package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type statuscode struct {
	Description string `json:"Description"`
	Code        int    `json:"Code"`
}

type codes []statuscode

func main() {
	var data codes

	rcvd := `[{"Description":"StatusOK","Code":200},{"Description":"StatusMovedPermanently","Code":301},{"Description":"StatusFound","Code":302},{"Description":"StatusSeeOther","Code":303},{"Description":"StatusTemporaryRedirect","Code":307},{"Description":"StatusBadRequest","Code":400},{"Description":"StatusUnauthorized","Code":401},{"Description":"StatusPaymentRequired","Code":402},{"Description":"StatusForbidden","Code":403},{"Description":"StatusNotFound","Code":404},{"Description":"StatusMethodNotAllowed","Code":405},{"Description":"StatusTeapot","Code":418},{"Description":"StatusInternalServerError","Code":500}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}
