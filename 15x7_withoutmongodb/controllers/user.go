package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/juandreww/2k22go/15x7_withoutmongodb/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var ctx = context.Background()

var maps = make(map[models.Contacts]models.Contacts)

type UserController struct {
}

func NewUserController(cl *mongo.Database) *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	data, err := uc.cl.Collection("contacts").Find(ctx, bson.M{"Fname": p.ByName("fname")})
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
		for _, v := range result {
			fmt.Println(v.Fname, v.Lname)
		}
	} else {
		fmt.Println("Not Found with Contact Name: " + p.ByName("fname"))
	}
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.Contacts{}
	json.NewDecoder(req.Body).Decode(&u)

	// u.Id = uuid.New().String()
	_, err := uc.cl.Collection("contacts").InsertOne(ctx, u)
	checkError(err)

	js, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", js)
}

func (uc UserController) UpdateUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.Contacts{}
	json.NewDecoder(req.Body).Decode(&u)

	selector := bson.M{"Fname": p.ByName("fname")}
	_, err := uc.cl.Collection("contacts").UpdateOne(ctx, selector, bson.M{"$set": u})
	checkError(err)

	fmt.Println("Update finished")
}

func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	data, err := uc.cl.Collection("contacts").Find(ctx, bson.M{"Fname": p.ByName("fname")})
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
		for _, v := range result {
			fmt.Println(v.Fname, v.Lname)
		}
	} else {
		fmt.Println("Not Found with Contact Name: " + p.ByName("fname"))
	}

	fname := p.ByName("fname")
	var selector = bson.M{"Fname": fname}
	_, err = uc.cl.Collection("contacts").DeleteOne(ctx, selector)
	checkError(err)

	data, err = uc.cl.Collection("contacts").Find(ctx, bson.M{"Fname": p.ByName("fname")})
	checkError(err)
	defer data.Close(ctx)

	result = make([]models.Contacts, 0)
	for data.Next(ctx) {
		var row models.Contacts
		err := data.Decode(&row)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Remaining Contacts"+strconv.Itoa(len(result))+"\n")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Remove contacts "+fname+" success\n")
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal("error lah")
		// log.Fatal(err)
	}
}
