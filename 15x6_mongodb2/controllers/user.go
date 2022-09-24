package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/juandreww/2k22go/15x6_mongodb2/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type UserController struct {
	cl *mongo.Database
}

func NewUserController(cl *mongo.Database) *UserController {
	return &UserController{cl}
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Println(p.ByName("fname"))
	u := models.Contacts{
		Fname: p.ByName("fname"),
		Lname: "Ayu",
	}

	data, err := uc.cl.Collection("contacts").Find(ctx, u)
	checkError(err)
	defer data.Close(ctx)

	result := make([]models.Contacts, 0)
	for data.Next(ctx) {
		var row models.Contacts
		err := data.Decode(&row)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("First Name : ", result[0].Fname)
		fmt.Println("Last Name : ", result[0].Lname)
	} else {
		fmt.Println("Not Found")
		fmt.Println(result)
	}
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.NewUser{}
	json.NewDecoder(req.Body).Decode(&u)

	// u.Id = uuid.New().String()
	_, err := uc.cl.Collection("contacts").InsertOne(context.TODO(), models.Contacts{"Hageko", "Batam"})
	checkError(err)

	js, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", js)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}

func checkError(err error) {
	if err != nil {
		log.Fatal("error lah")
		// log.Fatal(err)
	}
}
