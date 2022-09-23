package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/juandreww/2k22go/15x5_mongodb1/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

type UserController struct {
	cl *mongo.Collection
}

func NewUserController(cl *mongo.Collection) *UserController {
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

	u.Id = uuid.New().String()
	_, err := uc.cl.InsertOne(context.TODO(), student{"Wick", 2})
	if err != nil {
		log.Fatal(err)
	}

	js, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", js)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
