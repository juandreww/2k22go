package model

import (
	"2k22go/views"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func ReadAll() ([]views.Kelapa, error) {
	rows, err := con.Query("SELECT * FROM trnkelapabakar")
	if err != nil {
		return nil, err
	}
	
	coconut := []views.Kelapa{}
	spew.Dump(coconut)
	for rows.Next() {
		data := views.Kelapa{}
		rows.Scan(&data.Type2, &data.Quantity)
		coconut = append(coconut, data)
	}
	
	fmt.Println("CreateKelapa here...")
	// if err != nil {
	// 	return err
	// }
	// defer insertQ.Close()
	
	return coconut, nil
}