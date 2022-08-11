package model

import (
	"fmt"
)

func CreateKelapa() error {
	insertQ, err := con.Query("INSERT INTO trnkelapabakar VALUES(?, ?)", "Try", "Try")
	defer insertQ.Close()
	if err != nil {
		return err
	}
	fmt.Println("CreateKelapa here...")
	return nil
}