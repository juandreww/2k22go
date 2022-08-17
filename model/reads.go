package model

import (
	"2k22go/views"
	"fmt"
	"log"
	"github.com/davecgh/go-spew/spew"
	// "database/sql"
)

func ReadAll() ([]views.Kelapa, error) {
	rows, err := con.Query("SELECT type2, quantity FROM trnkelapabakar")
	if err != nil {
		fmt.Println("ya")
		log.Fatal(err)
	}
	defer rows.Close()
	
	coconut := []views.Kelapa{}
	i := 0;
	fmt.Println("wdaw")
	for rows.Next() {
		kelapa := views.Kelapa{}
		if i == 5 {
			break;
		}
		
		if err := rows.Scan(&kelapa.Type2, &kelapa.Quantity); err != nil {
            log.Fatal(err)
        }

		fmt.Printf("hey %s you %.2f\n", kelapa.Type2, kelapa.Quantity)
		i++
	}
	
	fmt.Println("CreateKelapa here...")
	
	return coconut, nil
}

func ReadSelected(uid string) ([]views.Kelapa, error) {
	rows, err := con.Query("SELECT * FROM trnkelapabakar WHERE uid = ?", uid)
	

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