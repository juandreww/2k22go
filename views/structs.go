package views

type Response struct {
	Code int `json: "code"`
	Body interface{} `json: "body"`
}

type Kelapa struct {
	type2 string `json: "type2"`
	quantity float64 `json: "quantity"`
}