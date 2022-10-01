package staffs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/juandreww/2k22go/18x1_gomongopostgres2/config"
	"gopkg.in/mgo.v2/bson"
)

var ctx = context.Background()

type Staff struct {
	ID       string
	Name     string
	UserName string
	Password string
	IsActive bool
}

func AllStaffs() ([]Staff, error) {
	rows, err := config.DB.Collection("employees").Find(ctx, bson.M{})
	checkError(err)
	defer rows.Close(ctx)

	list := make([]Staff, 0)
	for rows.Next(ctx) {
		p := Staff{}
		err := rows.Decode(&p)
		checkError(err)

		list = append(list, p)
	}

	return list, nil
}

func CreateStaff(r *http.Request) (Staff, error) {
	// get form values
	p := Staff{}
	p.Name = r.FormValue("name")
	p.UserName = r.FormValue("username")
	p.Password = r.FormValue("password")
	checkActive := r.FormValue("isactive")
	if checkActive == "on" {
		p.IsActive = true
	} else {
		p.IsActive = false
	}

	p.ID = uuid.New().String()

	// validate form values
	if p.Name == "" || p.UserName == "" || p.Password == "" {
		return p, errors.New("400. Bad request. All fields must be complete")
	}

	// insert values
	_, err := config.DB.Collection("employees").InsertOne(ctx, p)
	if err != nil {
		return p, errors.New("500. Internal Server Error." + err.Error())
	}
	return p, nil
}

func OneStaff(r *http.Request) (Staff, error) {
	p := Staff{}
	id := r.FormValue("id")
	if id == "" {
		return p, errors.New("400. Bad Request")
	}

	row, err := config.DB.Collection("employees").Find(ctx, bson.M{"id": id})
	fmt.Println(row)
	checkError(err)
	defer row.Close(ctx)

	for row.Next(ctx) {
		tmp := Staff{}
		err := row.Decode(&tmp)
		if err != nil {
			log.Fatal(err)
		}

		p = tmp
	}

	fmt.Println(p)

	return p, nil
}

func UpdateStaff(r *http.Request) (Staff, error) {
	// get form values
	p := Staff{}
	p.ID = r.FormValue("id")
	p.Name = r.FormValue("name")
	p.UserName = r.FormValue("username")
	p.Password = r.FormValue("password")

	checkActive := r.FormValue("isactive")
	if checkActive == "on" {
		p.IsActive = true
	} else {
		p.IsActive = false
	}

	// validate form values
	if p.Name == "" || p.UserName == "" || p.Password == "" {
		return p, errors.New("400. Bad request. All fields must be complete")
	}

	// insert values
	selector := bson.M{"id": p.ID}
	_, err := config.DB.Collection("employees").UpdateOne(ctx, selector, bson.M{"$set": p})
	if err != nil {
		return p, err
	}

	return p, nil
}

// func ModelDeleteStaff(r *http.Request) (Staff, error) {
// 	id := r.FormValue("id")
// 	if id == "" {
// 		return Staff{}, errors.New("400. Bad Request")
// 	}

// 	p, _ := OneStaff(r)

// 	_, err := config.DB.Exec("DELETE FROM employees WHERE id=$1;", id)
// 	if err != nil {
// 		return p, errors.New("500. Internal Server Error")
// 	}
// 	return p, nil
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
