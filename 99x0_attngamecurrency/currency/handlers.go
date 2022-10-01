package currency

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/juandreww/2k22go/99x0_attngamecurrency/config"
)

func Index(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "index.gohtml", nil)
}

func SaveCurrency(w http.ResponseWriter, r *http.Request) {
	data := ModelSaveCurrency(r)
	config.TPL.ExecuteTemplate(w, "index.gohtml", data)
}

func ListCurrency(w http.ResponseWriter, r *http.Request) {
	data := ModelListCurrency()

	config.TPL.ExecuteTemplate(w, "listcurrency.gohtml", data)
}

func ListCurrencyRate(w http.ResponseWriter, r *http.Request) {
	data := ModelListCurrencyRate(w)

	config.TPL.ExecuteTemplate(w, "listcurrencyrate.gohtml", data)
}

func AddConversionRate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("This is add conversion rate api: ", r.Method)
		data := ConfigConvertRate{
			r.FormValue("currencyfrom"),
			r.FormValue("currencyto"),
			r.FormValue("rate"),
		}

		var str string
		var intval int
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
			config.TPL.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
			return
		} else {
			intval, err = strconv.Atoi(check1.ID)
			if intval < 2 {
				tmp := Currency{
					"error",
					"One of the CurrencyID is not found in database",
				}
				config.TPL.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
				return
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
					"CurrencyRate is not exist in the database",
				}
				config.TPL.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
				return
			}
		} else {
			intval, err = strconv.Atoi(str)
			if intval > 0 {
				tmp := Currency{
					"error",
					"CurrencyRate already exist in the database",
				}
				config.TPL.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
				return
			}
		}

		sqlStatement = `SELECT nullif(max(id),0) id FROM currencyrate`
		row = config.DB.QueryRow(sqlStatement)
		err = row.Scan(&str)
		IsError = HandleErrorOfSelect(w, err)
		if IsError == true {
			str = "0"
		}

		intval, err = strconv.Atoi(str)
		intval++
		sqlStatement = `
			INSERT INTO currencyrate (id, currencyfrom, currencyto, rate)
			VALUES ($1, $2, $3, $4)`
		_, err = config.DB.Exec(sqlStatement, intval, data.CurrencyFrom, data.CurrencyTo, data.Rate)
		if err != nil {
			panic(err)
		}

		tmp := Currency{
			"succeed",
			"Currency Rate added",
		}
		config.TPL.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
	} else {
		config.TPL.ExecuteTemplate(w, "addconversionrate.gohtml", nil)
	}
}

func ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data := ModelConvertCurrency(w, r)
		config.TPL.ExecuteTemplate(w, "convertcurrency.gohtml", data)
	} else {
		config.TPL.ExecuteTemplate(w, "convertcurrency.gohtml", nil)
	}
}

func HandleErrorOfSelect(w http.ResponseWriter, err error) bool {
	data := false
	switch err {
	case sql.ErrNoRows:
		data = true
	case nil:
		data = false
	default:
		data = true
	}

	return data
}
