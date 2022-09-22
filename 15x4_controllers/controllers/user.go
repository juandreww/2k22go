package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/juandreww/2k22go/15x3_postdelete1/models"
	"github.com/julienschmidt/httprouter"
)

func GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
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

func CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.NewUser{}

	json.NewDecoder(req.Body).Decode(&u)

	u.Id = uuid.New().String()
	js, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", js)
}

func DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
