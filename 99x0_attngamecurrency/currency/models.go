package currency

type Currency struct {
	ID   string
	Name string
}

type ConfigConvertRate struct {
	CurrencyFrom string
	CurrencyTo   string
	Rate         string
}
