package models

import (
	"api_assessment/datab"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Login struct {
	Username string
	Password string
}

type Validation struct {
	Value string
	Valid string
}

func (user *User) Validate() *RestErr {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return NewBadRequestError("Invalid Email Address")
	}
	if user.Password == "" {
		return NewBadRequestError("Invalid Password")
	}
	return nil
}

func (user *User) Save() *RestErr {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(*datab.GiveSchema())

	defer db.Close()

	insert, err := db.Queryx(
		"INSERT INTO users (id, name, email, password, role) VALUES (($1),($2),($3), ($4), ($5))",
		user.ID, user.Name, user.Email, user.Password, "PATIENT")

	if insert == nil {

	}

	return nil

}

func (user *User) GetByEmail() *RestErr {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(*datab.GiveSchema())

	defer db.Close()

	results, err := db.Queryx("SELECT * FROM userrs where email=($1)", user)
	if err != nil {
		return NewInternalServerError("Invalid Email")
	}

	results.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

}
