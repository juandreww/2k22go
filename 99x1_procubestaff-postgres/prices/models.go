package prices

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/juandreww/2k22go/18x0_gomongopostgres/config"
)

type Pricing struct {
	ID    string
	Title string
	Price float32
}

func AllPrices() ([]Pricing, error) {
	rows, err := config.DB.Query("SELECT * FROM pricing;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prices := make([]Pricing, 0)
	for rows.Next() {
		pc := Pricing{}
		err := rows.Scan(&pc.ID, &pc.Title, &pc.Price)
		if err != nil {
			return nil, err
		}
		prices = append(prices, pc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return prices, nil
}

func OnePrice(r *http.Request) (Pricing, error) {
	pc := Pricing{}
	id := r.FormValue("id")
	if id == "" {
		return pc, errors.New("400. Bad Request")
	}

	row := config.DB.QueryRow("SELECT * FROM pricing WHERE id = $1;", id)
	err := row.Scan(&pc.ID, &pc.Title, &pc.Price)

	if err != nil {
		return pc, err
	}

	return pc, nil
}

func PutPrice(r *http.Request) (Pricing, error) {
	// get form values
	bk := Pricing{}
	bk.ID = r.FormValue("id")
	bk.Title = r.FormValue("title")
	p := r.FormValue("price")

	// validate form values
	if bk.ID == "" || bk.Title == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number.")
	}
	bk.Price = float32(f64)

	// insert values
	_, err = config.DB.Exec("INSERT INTO pricing (id, title, price) VALUES ($1, $2, $3)", bk.ID, bk.Title, bk.Price)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func UpdatePrice(r *http.Request) (Pricing, error) {
	// get form values
	bk := Pricing{}
	bk.ID = r.FormValue("id")
	bk.Title = r.FormValue("title")
	p := r.FormValue("price")

	// validate form values
	if bk.ID == "" || bk.Title == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty.")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("400. Bad Request. Fields can't be empty.")
	}
	bk.Price = float32(f64)

	// insert values
	_, err = config.DB.Exec("UPDATE pricing SET id = $1, title=$2, price=$3 WHERE id=$1;", bk.ID, bk.Title, bk.Price)
	if err != nil {
		return bk, err
	}

	return bk, nil
}

func DeletePrice(r *http.Request) error {
	id := r.FormValue("id")
	if id == "" {
		return errors.New("400. Bad Request.")
	}

	_, err := config.DB.Exec("DELETE FROM pricing WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
