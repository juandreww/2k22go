package staffs

import (
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
