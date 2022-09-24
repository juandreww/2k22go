package models

type NewUser struct {
	Name       string `json:"Name" bson:"name"`
	Email      string `json:"Email" bson:"gender"`
	Occupation string `json:"Occupation" bson:"age"`
	Location   string `json:"Location" bson:"location"`
	Age        string `json:"Age" bson:"age"`
	Id         string `json:"id" bson:"_id"`
}

type Contacts struct {
	Fname string `bson:"Fname"`
	Lname string `bson:"Lname"`
}
