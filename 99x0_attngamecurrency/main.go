package main

import (
	"net/http"

	"github.com/juandreww/2k22go/99x0_attngamecurrency/currency"
	_ "github.com/lib/pq"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", currency.Index)
	mux.HandleFunc("/savecurrency", currency.SaveCurrency)
	mux.HandleFunc("/listcurrency", currency.ListCurrency)
	mux.HandleFunc("/listcurrencyrate", currency.ListCurrencyRate)
	mux.HandleFunc("/addconversionrate", currency.AddConversionRate)
	mux.HandleFunc("/convertcurrency", currency.ConvertCurrency)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
