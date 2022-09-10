package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	// "os"
)

const (
	host = "ec2-52-73-155-171.compute-1.amazonaws.com"
	port = 5432
	user = "'awtpbzyctlpydm'"
	password = "a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96"
	dbname = "d4ehughfapgq0k"
)

type currency struct {
	ID string
	Name string
}

type configconvertrate struct {
	CurrencyFrom string
	CurrencyTo string
	Rate string
}

var tpl *template.Template
var con *sql.DB

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=require",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	con = db
	return db
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/savecurrency", savecurrency)
	mux.HandleFunc("/listcurrency", listcurrency)
	mux.HandleFunc("/listcurrencyrate", listcurrencyrate)
	mux.HandleFunc("/addconversionrate", addconversionrate)
	mux.HandleFunc("/convertcurrency", convertcurrency)
	db := ConnectDB()
	defer db.Close()

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is index api: ", r.Method)

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func savecurrency(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is savecurrency api: ", r.Method)

	data := currency{
		r.FormValue("id"),
		r.FormValue("name"),
	}

	fmt.Println(data)

	value := currency{}
	isexist := false
	sqlStatement := `SELECT id, name FROM currency WHERE (id=$1 OR name=$2);`
	row := con.QueryRow(sqlStatement, data.ID, data.Name)
	err := row.Scan(&value.ID, &value.Name,)
	switch err {
	case sql.ErrNoRows:
		isexist = false
	case nil:
		isexist = true
	default:
	  	panic(err)
	}
	
	if isexist == false {
		sqlStatement = `
			INSERT INTO currency (id, name)
			VALUES ($1, $2)`
		_, err := con.Exec(sqlStatement, data.ID, data.Name)
		if err != nil {
			panic(err)
		}
	} else {
		data = currency{
			"error",
			"error",
		}
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func listcurrency(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is listcurrency api: ", r.Method)
	var list []currency

	rows, err := con.Query("SELECT id, name FROM currency ORDER BY id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		value := currency{}
		err := rows.Scan(&value.ID, &value.Name,)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("row is not exist")
			return
		case nil:
			fmt.Println(value.ID, value.Name)
		default:
			panic(err)
		}

		list = append(list, value)
	}

	tpl.ExecuteTemplate(w, "listcurrency.gohtml", list)
}

func listcurrencyrate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is listcurrencyrate api: ", r.Method)
	var list []configconvertrate

	rows, err := con.Query("SELECT cf.name currencyfrom, ct.name currencyto, round(p.rate,2) rate FROM currencyrate p LEFT JOIN currency cf ON cf.id = p.currencyfrom LEFT JOIN currency ct ON ct.id = p.currencyto ORDER BY p.id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		value := configconvertrate{}
		err := rows.Scan(&value.CurrencyFrom, &value.CurrencyTo, &value.Rate)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("row is not exist")
			return
		case nil:
			fmt.Println(value.CurrencyFrom, value.CurrencyTo, value.Rate)
		default:
			panic(err)
		}

		list = append(list, value)
	}

	tpl.ExecuteTemplate(w, "listcurrencyrate.gohtml", list)
}

func addconversionrate(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPost) {
		fmt.Println("This is add conversion rate api: ", r.Method)
		data := configconvertrate{
			r.FormValue("currencyfrom"),
			r.FormValue("currencyto"),
			r.FormValue("rate"),
		}

		var value string
		var intval int
		check1 := currency{}

		sqlStatement := `SELECT count(id) id FROM currency WHERE (id=$1 OR id=$2);`
		row := con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo)
		err := row.Scan(&check1.ID)
		isError := HandleErrorOfSelect(w, err)
		if (isError == true) {
			tmp := currency{
				"error",
				"All CurrencyID is not found in database",
			}
			tpl.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
			return
		} else {
			intval, err = strconv.Atoi(check1.ID)
			if intval < 2 {
				tmp := currency{
					"error",
					"One of the CurrencyID is not found in database",
				}
				tpl.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
				return
			}
		}
		
		sqlStatement = `SELECT count(id) id FROM currencyrate 
						WHERE ((currencyfrom=$1 AND currencyto=$2) OR (currencyfrom=$3 AND currencyto=$4))`
		row = con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo, data.CurrencyTo, data.CurrencyFrom)
		err = row.Scan(&value)
		isError = HandleErrorOfSelect(w, err)
		if (isError == true) {
			if value != "0" {
				tmp := currency{
					"error",
					"CurrencyRate is not exist in the database",
				}
				tpl.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
				return
			}
		} else {
			intval, err = strconv.Atoi(value)
			if intval >  0 {
				tmp := currency{
					"error",
					"CurrencyRate already exist in the database",
				}
				tpl.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
				return
			}
		}

		sqlStatement = `SELECT nullif(max(id),0) id FROM currencyrate`
		row = con.QueryRow(sqlStatement)
		err = row.Scan(&value)
		isError = HandleErrorOfSelect(w, err)
		if (isError == true) {
			value = "0"
		}

		intval, err = strconv.Atoi(value)
		intval++
		sqlStatement = `
			INSERT INTO currencyrate (id, currencyfrom, currencyto, rate)
			VALUES ($1, $2, $3, $4)`
		_, err = con.Exec(sqlStatement, intval, data.CurrencyFrom, data.CurrencyTo, data.Rate)
		if err != nil {
			panic(err)
		}
		
		tmp := currency{
			"succeed",
			"Currency Rate added",
		}
		tpl.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
	} else {
		fmt.Println("This is add conversion rate api: ", r.Method)

		tpl.ExecuteTemplate(w, "addconversionrate.gohtml", nil)
	}
}

func convertcurrency(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPost) {
		fmt.Println("This is add conversion rate api: ", r.Method)
		data := configconvertrate{
			r.FormValue("currencyfrom"),
			r.FormValue("currencyto"),
			r.FormValue("amount"),
		}
		fmt.Println(data)

		var value string
		var intval int
		var floatval float64
		var amount float64
		check1 := currency{}

		sqlStatement := `SELECT count(id) id FROM currency WHERE (id=$1 OR id=$2);`
		row := con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo)
		err := row.Scan(&check1.ID)
		isError := HandleErrorOfSelect(w, err)
		if (isError == true) {
			tmp := currency{
				"error",
				"All CurrencyID is not found in database",
			}
			tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
			return
		} else {
			intval, err = strconv.Atoi(check1.ID)
			if intval < 2 {
				tmp := currency{
					"error",
					"One of the CurrencyID is not found in database",
				}
				tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
				return
			}
		}
		
		sqlStatement = `SELECT count(id) id FROM currencyrate 
						WHERE ((currencyfrom=$1 AND currencyto=$2) OR (currencyfrom=$3 AND currencyto=$4))`
		row = con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo, data.CurrencyTo, data.CurrencyFrom)
		err = row.Scan(&value)
		isError = HandleErrorOfSelect(w, err)
		if (isError == true) {
			if value != "0" {
				tmp := currency{
					"error",
					"CurrencyRate is not exist in the database (1)",
				}
				tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
				return
			}
		}

		var val1, val2 string
		sqlStatement = `SELECT currencyfrom, currencyto,rate FROM currencyrate 
					WHERE ((currencyfrom=$1 AND currencyto=$2) OR (currencyfrom=$3 AND currencyto=$4))
					LIMIT 1`
		row = con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo, data.CurrencyTo, data.CurrencyFrom)
		err = row.Scan(&val1, &val2, &value)
		isError = HandleErrorOfSelect(w, err)
		if (isError == true) {
			tmp := currency{
				"error",
				"CurrencyRate is not exist in the database (2)",
			}
			tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
			return 
		}

		floatval, err = strconv.ParseFloat(value, 64)
		amount, err = strconv.ParseFloat(data.Rate, 64)
		fmt.Println("value is",floatval * amount)
		
		// tmp := currency{
		// 	"succeed",
		// 	"Currency Rate added",
		// }
		// tpl.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
	} else {
		fmt.Println("This is add convertcurrency api: ", r.Method)

		tpl.ExecuteTemplate(w, "convertcurrency.gohtml", nil)
	}
}

func HandleErrorOfSelect(w http.ResponseWriter, err error) bool {
	data := false
	switch err {
	case sql.ErrNoRows:
		fmt.Println(11)
		data = true
	case nil:
		fmt.Println(22)
		data = false
	default:
		fmt.Println(33)
		data = true
	}

	return data
}