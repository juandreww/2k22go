package model

import (
	"2k22go/views"
	"fmt"
	"log"
	"github.com/davecgh/go-spew/spew"
)

func ReadAll() ([]views.Kelapa, error) {
	rows, err := con.Query("SELECT type2 FROM trnkelapabakar")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	coconut := []views.Kelapa{}
	i := 0;
	fmt.Println("wdaw")
	for rows.Next() {
		// var Type2 string
		
		kelapa := views.Kelapa{}
		if i == 5 {
			break;
		}
		
		if err := rows.Scan(&kelapa.Type2); err != nil {
            log.Fatal(err)
        }

		fmt.Printf("hey %s you\n", kelapa.Type2)
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