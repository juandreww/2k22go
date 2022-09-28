package models

import (
	"errors"
	"net/http"
)

type Pricing struct {
	ID    string
	Title string
	Price float32
}

func OnePrice(r *http.Request) (Pricing, error) {
	pc := Pricing{}
	id := r.FormValue("id")
	if id == "" {
		return pc, errors.New("400. Bad Request")
	}

	row := db.QueryRow("SELECT * FROM pricing WHERE id = $1;", id)
	err := row.Scan(&pc.ID, &pc.Title, &pc.Price)

	if err != nil {
		return pc, err
	}

	return pc, nil
}
