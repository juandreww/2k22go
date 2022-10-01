package staffs

import (
	"database/sql"
	"net/http"

	"github.com/juandreww/2k22go/18x1_gomongopostgres2/config"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	list, err := AllStaffs()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	config.TPL.ExecuteTemplate(w, "staffs.gohtml", list)
}

func NewStaffForm(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "newstaff.gohtml", nil)
}

func NewStaffSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	bk, err := CreateStaff(r)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	// confirm insertion
	config.TPL.ExecuteTemplate(w, "newstaffcreated.gohtml", bk)
}

func EditStaffForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	p, err := OneStaff(r)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "editstaffform.gohtml", p)
}

func EditStaffSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	p, err := UpdateStaff(r)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	config.TPL.ExecuteTemplate(w, "editstaffsave.gohtml", p)
}

func DeleteStaff(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	p, err := ModelDeleteStaff(r)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "deletestaff.gohtml", p)
}
