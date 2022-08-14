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
	i := 0;
	for rows.Next() {
		if i == 5 {
			break;
		}
		data := views.Kelapa{}
		if i == 1 {
			fmt.Println(rows)
			fmt.Println(rows.Scan(&data.Type2, &data.Quantity))
		}
		rows.Scan(&data.Type2, &data.Quantity)
		coconut = append(coconut, data)
		i++
	}
	
	fmt.Println("CreateKelapa here...")
	
	return coconut, nil
}

func ReadSelected(uid string) ([]views.Kelapa, error) {
	spew.Dump(uid)
	rows, err := con.Query("SELECT * FROM trnkelapabakar WHERE type2 = ?", uid)
	spew.Dump(rows)

	if rows == nil {
		fmt.Println("No rows returned")
	}
	if err != nil {
		return nil, err
	} 
	
	coconut := []views.Kelapa{}
	// spew.Dump(coconut)
	
	for rows.Next() {
		data := views.Kelapa{}
		
		rows.Scan(&data.Type2, &data.Quantity)
		coconut = append(coconut, data)
	}
	
	fmt.Println("CreateKelapa here...")
	
	return coconut, nil
}