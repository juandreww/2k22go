package staffs

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/juandreww/2k22go/18x0_gomongopostgres/config"
)

type Staff struct {
	ID       string
	Name     string
	UserName string
	Password string
	IsActive bool
}

func AllStaffs() ([]Staff, error) {
	rows, err := config.DB.Query("SELECT * FROM employees;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]Staff, 0)
	for rows.Next() {
		p := Staff{}
		err := rows.Scan(&p.ID, &p.Name, &p.UserName, &p.Password, &p.IsActive)
		checkError(err)

		list = append(list, p)
	}

	return list, nil
}

func CreateStaff(r *http.Request) (Staff, error) {
	// get form values
	p := Staff{}
	p.Name = r.FormValue("name")
	p.UserName = r.FormValue("username")
	p.Password = r.FormValue("password")
	checkActive := r.FormValue("isactive")
	if checkActive == "on" {
		p.IsActive = true
	} else {
		p.IsActive = false
	}

	p.ID = uuid.New().String()

	// validate form values
	if p.Name == "" || p.UserName == "" || p.Password == "" {
		return p, errors.New("400. Bad request. All fields must be complete")
	}

	// insert values
	_, err := config.DB.Exec("INSERT INTO employees (id, name, username, password, isactive) VALUES ($1, $2, $3, $4, $5)", p.ID, p.Name, p.UserName, p.Password, p.IsActive)
	if err != nil {
		return p, errors.New("500. Internal Server Error." + err.Error())
	}
	return p, nil
}

func OneStaff(r *http.Request) (Staff, error) {
	p := Staff{}
	id := r.FormValue("id")
	if id == "" {
		return p, errors.New("400. Bad Request")
	}

	row := config.DB.QueryRow("SELECT * FROM employees WHERE id = $1;", id)
	err := row.Scan(&p.ID, &p.Name, &p.UserName, &p.Password, &p.IsActive)
	if err != nil {
		return p, err
	}

	return p, nil
}

func UpdateStaff(r *http.Request) (Staff, error) {
	// get form values
	p := Staff{}
	p.ID = r.FormValue("id")
	p.Name = r.FormValue("name")
	p.UserName = r.FormValue("username")
	p.Password = r.FormValue("password")

	checkActive := r.FormValue("isactive")
	if checkActive == "on" {
		p.IsActive = true
	} else {
		p.IsActive = false
	}

	// validate form values
	if p.Name == "" || p.UserName == "" || p.Password == "" {
		return p, errors.New("400. Bad request. All fields must be complete")
	}

	// insert values
	_, err := config.DB.Exec("UPDATE employees SET name = $1, username=$2, password = $3, isactive = $4 WHERE id=$5;", p.Name, p.UserName, p.Password, p.IsActive, p.ID)
	if err != nil {
		return p, err
	}

	return p, nil
}

func ModelDeleteStaff(r *http.Request) (Staff, error) {
	id := r.FormValue("id")
	if id == "" {
		return Staff{}, errors.New("400. Bad Request.")
	}

	p, err := OneStaff(r)

	_, err = config.DB.Exec("DELETE FROM employees WHERE id=$1;", id)
	if err != nil {
		return p, errors.New("500. Internal Server Error")
	}
	return p, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
