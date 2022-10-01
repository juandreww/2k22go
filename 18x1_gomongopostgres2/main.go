package main

import (
	"fmt"
	"net/http"

	"github.com/juandreww/2k22go/18x1_gomongopostgres2/staffs"
)

func main() {
	http.HandleFunc("/staff", staffs.List)
	http.HandleFunc("/staff/new/form", staffs.NewStaffForm)
	http.HandleFunc("/staff/new/save", staffs.NewStaffSave)
	http.HandleFunc("/staff/edit/form", staffs.EditStaffForm)
	http.HandleFunc("/staff/edit/save", staffs.EditStaffSave)
	http.HandleFunc("/staff/delete", staffs.DeleteStaff)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I love food")
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
