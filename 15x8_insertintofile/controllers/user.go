package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/juandreww/2k22go/15x8_insertintofile/models"
	"github.com/julienschmidt/httprouter"
)

// var ctx = context.Background()

type UserController struct {
	session map[string]models.Contacts
}

func NewUserController(mp map[string]models.Contacts) *UserController {
	return &UserController{mp}
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fname := p.ByName("fname")
	ct := uc.session[fname]

	ctm, err := json.Marshal(ct)
	checkError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", ctm)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fn, err := os.Create("data.txt")
	checkError(err)
	defer fn.Close()

	u := models.Contacts{}
	json.NewDecoder(req.Body).Decode(&u)
	fname := u.Fname

	// u.Id = uuid.New().String()
	uc.session[fname] = u

	fmt.Println(uc)
	js, _ := json.Marshal(u)
	json.NewEncoder(fn).Encode(uc.session)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", js)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fn, err := os.Create("data.txt")
	checkError(err)
	defer fn.Close()

	delete(uc.session, p.ByName("fname"))
	json.NewEncoder(fn).Encode(uc.session)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Deleted record of %s\n", p.ByName("fname"))
}

func checkError(err error) {
	if err != nil {
		log.Fatal("error lah")
		// log.Fatal(err)
	}
}
