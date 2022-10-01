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

// func CreateStaff(r *http.Request) (Staff, error) {
// 	// get form values
// 	bk := Staff{}
// 	bk.ID = r.FormValue("id")
// 	bk.Title = r.FormValue("title")
// 	p := r.FormValue("price")

// 	// validate form values
// 	if bk.ID == "" || bk.Title == "" || p == "" {
// 		return bk, errors.New("400. Bad request. All fields must be complete.")
// 	}

// 	// convert form values
// 	f64, err := strconv.ParseFloat(p, 32)
// 	if err != nil {
// 		return bk, errors.New("406. Not Acceptable. Price must be a number.")
// 	}
// 	bk.Price = float32(f64)

// 	// insert values
// 	_, err = config.DB.Exec("INSERT INTO pricing (id, title, price) VALUES ($1, $2, $3)", bk.ID, bk.Title, bk.Price)
// 	if err != nil {
// 		return bk, errors.New("500. Internal Server Error." + err.Error())
// 	}
// 	return bk, nil
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
