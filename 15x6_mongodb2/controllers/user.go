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

type UserController struct {
	cl *mongo.Database
}

func NewUserController(cl *mongo.Database) *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.NewUser{
		Name:       "Garry Chapman",
		Email:      "garrychapmanbros@gmail.com",
		Occupation: "Account Executive",
		Location:   "South UK",
		Age:        "35",
		Id:         p.ByName("id"),
	}

	js, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", js)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.NewUser{}

	json.NewDecoder(req.Body).Decode(&u)

	fmt.Println("aahaaaaaaaaaa")
	// u.Id = uuid.New().String()
	_, err := uc.cl.Collection("contacts").InsertOne(context.TODO(), models.Contacts{"Hageko", "Batam"})
	checkError(err)

	js, _ := json.Marshal(u)
	fmt.Println("dadaaaaaaaaaa")
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
		log.Fatal(err)
	}
}
