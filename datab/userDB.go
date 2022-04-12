package datab

import (
	"api_assessment/models"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

//Function to add a doctor to the DB
func AddUsers(u models.User) {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	insert, err := db.Queryx(
		"INSERT INTO users (id, name, email, password, role) VALUES (($1),($2),($3),($4),($5))",
		u.ID, u.Name, u.Email, u.Password, u.Role)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

//Function for getting user by username
func GetUser(username string) *models.User {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	u := &models.User{}

	results, err := db.Queryx("SELECT * FROM users where name=($1)", username)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return u
}

func CreateUser(user models.User) (*models.User, *models.RestErr) {
	if err := Validate(&user); err != nil {
		return nil, err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, models.NewBadRequestError("Failed to Encrypt Password")
	}

	user.Password = string(pwSlice[:])

	if err := Save(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func TestGetUser(user models.User) (*models.User, *models.RestErr) {
	result := &models.User{Email: user.Email}

	if err := GetByEmail(result); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, models.NewBadRequestError("Failed to Decrypt Password")
	}

	resultWp := &models.User{ID: result.ID, Name: result.Name, Email: result.Email, Role: result.Role}
	return resultWp, nil
}

func Validate(user *models.User) *models.RestErr {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return models.NewBadRequestError("Invalid Email Address")
	}
	if user.Password == "" {
		return models.NewBadRequestError("Invalid Password")
	}
	return nil
}

func Save(user *models.User) *models.RestErr {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(*GiveSchema())

	log.Println("Testing that Save function is being accessed correctly:", user)

	insert, err := db.Queryx("INSERT INTO users (id, name, email, password, role) VALUES (($1),($2),($3),($4),($5))",
		user.ID, user.Name, user.Email, user.Password, user.Role)

	if insert == nil {
	}

	if err != nil {

	}
	defer db.Close()
	return nil

}

//Function to fetch a user by their email
func GetByEmail(user *models.User) *models.RestErr {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(*GiveSchema())

	defer db.Close()

	results, err := db.Queryx("SELECT * FROM userrs where email=($1)", user)
	if err != nil {
		return models.NewInternalServerError("Invalid Email")
	}

	results.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	return nil
}
