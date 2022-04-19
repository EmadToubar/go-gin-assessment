package datab

import (
	"api_assessment/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type userDB struct {
	DB *sqlx.DB
}

type UserDB interface {
	AddUsers(u models.User)
	GetUser(username string) (*models.User, error)
	CreateUser(user models.User) (*models.User, *models.RestErr)
	TestGetUser(user models.User) (*models.User, *models.RestErr)
	Validate(user *models.User) *models.RestErr
	Save(user *models.User) *models.RestErr
	GetByEmail(user *models.User) *models.RestErr
}

func UserDBProvider(ctx *sqlx.DB) UserDB {
	return &userDB{
		DB: ctx,
	}
}

//Function to add a doctor to the DB
func (us *userDB) AddUsers(u models.User) {
	if insert, err := us.DB.Queryx(
		"INSERT INTO users (id, name, email, password, role) VALUES (($1),($2),($3),($4),($5))",
		u.ID, u.Name, u.Email, u.Password, u.Role); err != nil {
		panic(err.Error())
	} else {
		insert.Close()
	}
}

//Function for getting user by username
func (us *userDB) GetUser(username string) (*models.User, error) {
	// db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable") this way you need to connect to your database only once
	u := &models.User{}
	if results, err := us.DB.Queryx("SELECT * FROM users where name=($1)", username); err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	} else {
		if results.Next() {
			if err = results.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role); err != nil {
				return nil, err
			}
		}
	}
	return u, nil
}

func (us *userDB) CreateUser(user models.User) (*models.User, *models.RestErr) {
	if err := us.Validate(&user); err != nil {
		return nil, err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, models.NewBadRequestError("Failed to Encrypt Password")
	}

	user.Password = string(pwSlice[:])

	if err := us.Save(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *userDB) TestGetUser(user models.User) (*models.User, *models.RestErr) {
	if result, err := us.GetUser(user.Name); err != nil {
		return nil, &models.RestErr{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
			Error:   "Error Retriving user " + user.Name,
		}
	} else {
		if err := us.GetByEmail(result); err != nil {
			return nil, err
		} else {
			if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
				return nil, models.NewBadRequestError("Failed to Decrypt Password")
			} else {
				resultWp := &models.User{ID: result.ID, Name: result.Name, Email: result.Email, Role: result.Role}
				return resultWp, nil
			}
		}
	}
}

func (us *userDB) Validate(user *models.User) *models.RestErr {
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

func (us *userDB) Save(user *models.User) *models.RestErr {
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
func (us *userDB) GetByEmail(user *models.User) *models.RestErr {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(*GiveSchema())

	defer db.Close()

	results, err := db.Queryx("SELECT * FROM users where email=($1)", user.Email)

	if err != nil {
		return models.NewInternalServerError("Invalid Email")
	}

	results.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	return nil
}
