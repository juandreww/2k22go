package currency

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/juandreww/2k22go/99x0_attngamecurrency/config"
)

type Currency struct {
	ID   string
	Name string
}

type ConfigConvertRate struct {
	CurrencyFrom string
	CurrencyTo   string
	Rate         string
}

func ModelSaveCurrency(r *http.Request) Currency {
	data := Currency{
		r.FormValue("id"),
		r.FormValue("name"),
	}

	cur := Currency{}
	IsExist := false
	sqlStatement := `SELECT id, name FROM currency WHERE (id=$1 OR name=$2);`
	row := config.DB.QueryRow(sqlStatement, data.ID, data.Name)
	err := row.Scan(&cur.ID, &cur.Name)
	switch err {
	case sql.ErrNoRows:
		IsExist = false
	case nil:
		IsExist = true
	default:
		panic(err)
	}

	if IsExist == false {
		sqlStatement = `
			INSERT INTO currency (id, name)
			VALUES ($1, $2)`
		_, err := config.DB.Exec(sqlStatement, data.ID, data.Name)
		if err != nil {
			panic(err)
		}
	} else {
		data = Currency{
			"error",
			"error",
		}
	}

	return data
}

func ModelListCurrency() []Currency {
	var list []Currency

	rows, err := config.DB.Query("SELECT id, name FROM currency ORDER BY id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		cur := Currency{}
		err := rows.Scan(&cur.ID, &cur.Name)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("row is not exist")
			return []Currency{}
		case nil:
			fmt.Println(cur.ID, cur.Name)
		default:
			panic(err)
		}

		list = append(list, cur)
	}
	return list
}

func ModelListCurrencyRate(w http.ResponseWriter) []ConfigConvertRate {
	var list []ConfigConvertRate

	rows, err := config.DB.Query("SELECT cf.name currencyfrom, ct.name currencyto, round(p.rate,2) rate FROM currencyrate p LEFT JOIN currency cf ON cf.id = p.currencyfrom LEFT JOIN currency ct ON ct.id = p.currencyto ORDER BY p.id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		cur := ConfigConvertRate{}
		err := rows.Scan(&cur.CurrencyFrom, &cur.CurrencyTo, &cur.Rate)

		IsError := HandleErrorOfSelect(w, err)
		if IsError == true {
			fmt.Println("row is not exist")
			return []ConfigConvertRate{}
		}

		list = append(list, cur)
	}
	return list
}
