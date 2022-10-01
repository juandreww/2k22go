package currency

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func ModelConvertCurrency(w http.ResponseWriter, r *http.Request) Currency {
	data := ConfigConvertRate{
		r.FormValue("currencyfrom"),
		r.FormValue("currencyto"),
		r.FormValue("amount"),
	}
	fmt.Println(data)

	var str string
	var intval int
	var floatval, amount float64
	check1 := Currency{}

	sqlStatement := `SELECT count(id) id FROM currency WHERE (id=$1 OR id=$2);`
	row := config.DB.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo)
	err := row.Scan(&check1.ID)
	IsError := HandleErrorOfSelect(w, err)
	if IsError == true {
		tmp := Currency{
			"error",
			"All CurrencyID is not found in database",
		}
		config.TPL.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
		return Currency{}
	} else {
		intval, err = strconv.Atoi(check1.ID)
		if intval < 2 {
			tmp := Currency{
				"error",
				"One of the CurrencyID is not found in database",
			}
			config.TPL.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
			return Currency{}
		}
	}

	sqlStatement = `SELECT count(id) id FROM currencyrate 
					WHERE ((currencyfrom=$1 AND currencyto=$2) OR (currencyfrom=$3 AND currencyto=$4))`
	row = config.DB.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo, data.CurrencyTo, data.CurrencyFrom)
	err = row.Scan(&str)
	IsError = HandleErrorOfSelect(w, err)
	if IsError == true {
		if str != "0" {
			tmp := Currency{
				"error",
				"CurrencyRate is not exist in the database (1)",
			}
			config.TPL.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
			return Currency{}
		}
	}

	var val1, val2 string
	sqlStatement = `SELECT currencyfrom, currencyto,rate FROM currencyrate 
				WHERE ((currencyfrom=$1 AND currencyto=$2) OR (currencyfrom=$3 AND currencyto=$4))
				LIMIT 1`
	row = config.DB.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo, data.CurrencyTo, data.CurrencyFrom)
	err = row.Scan(&val1, &val2, &str)
	IsError = HandleErrorOfSelect(w, err)
	if IsError == true {
		tmp := Currency{
			"error",
			"CurrencyRate is not exist in the database (2)",
		}
		config.TPL.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
		return Currency{}
	}

	floatval, err = strconv.ParseFloat(str, 64)
	amount, err = strconv.ParseFloat(data.Rate, 64)
	if data.CurrencyFrom == val1 {
		amount = amount * floatval
	} else {
		amount = amount / floatval
	}

	str = fmt.Sprintf("%.2f", amount)

	tmp := Currency{
		"succeed",
		"Congrats! You converted rate to " + str,
	}

	return tmp
}
