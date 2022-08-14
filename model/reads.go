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
	defer rows.Close()
	
	coconut := []views.Kelapa{}
	i := 0;
	for rows.Next() {
		if i == 5 {
			break;
		}
		var type2 string
		var quantity float64
		if i == 1 {
			// fmt.Println(rows)
		}
		rows.Scan(&type2, &quantity)
		fmt.Println(type2)
		// coconut = append(coconut, type2)
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