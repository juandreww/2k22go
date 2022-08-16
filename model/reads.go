package model

import (
	"2k22go/views"
	"fmt"
	"log"
	"github.com/davecgh/go-spew/spew"
)

func ReadAll() ([]views.Kelapa, error) {
	rows, err := con.Query("SELECT Quantity FROM trnkelapabakar")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	coconut := []views.Kelapa{}
	i := 0;
	fmt.Println("wdaw")
	for rows.Next() {
		var type2 string
		var Quantity float64
		if i == 5 {
			break;
		}
		
		if err := rows.Scan(&Quantity); err != nil {
            log.Fatal(err)
        }

		fmt.Printf("hey %s you %.2f", type2, Quantity)
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