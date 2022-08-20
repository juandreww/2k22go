package model

import (
	// "2k22go/views"
	"fmt"
	// "log"
	// "github.com/davecgh/go-spew/spew"
	// "database/sql"
)

func DeleteSelected(uid string) error {
	insertQ, err := con.Query("DELETE FROM trnkelapabakar WHERE uid = ($1)::uuid", uid)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
	// return 
}