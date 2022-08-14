package model

import (
	"fmt"
)

func CreateKelapa() error {
	insertQ, err := con.Query("INSERT INTO trnkelapabakar(type2, quantity) VALUES(?, ?)", "Try", "Try")
	
	fmt.Println("CreateKelapa here...")
	if err != nil {
		return err
	}
	defer insertQ.Close()
	
	return nil
}