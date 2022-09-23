package models

type NewUser struct {
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Occupation string `json:"Occupation"`
	Location   string `json:"Location"`
	Age        string `json:"Age"`
	Id         string `json:"id"`
}
