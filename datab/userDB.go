package datab

import (
	"api_assessment/helpers"
	"api_assessment/models"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		"INSERT INTO user (id, name, email, password, role) VALUES (($1),($2),($3),($4),($5))",
		u.ID, u.Name, u.Email, u.Password, u.Role)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func GetUser(username string) *models.User {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	u := &models.User{}

	results, err := db.Queryx("SELECT * FROM user where name=($1)", username)

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

func createAccount() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	users := [1]models.User{
		{Name: "Test", Email: "dummy@test.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Name))
		user := models.User{Name: users[i].Name, Email: users[i].Email, Password: generatedPassword}

		patient := models.Patient{ID: "50", Name: "Tester", Role: "PATIENT"}

		AddPatients(patient)
		AddUsers(user)

		// insert, err := db.Queryx(
		// 	"INSERT INTO patient (id, name, role) VALUES (($1),($2),($3))",
		// 	patient.ID, patient.Name, patient.Role)

		// // if there is an error inserting, handle it
		// if err != nil {
		// 	panic(err.Error())
		// }

		//defer insert.Close()
	}

}

func userLogin(username string, pass string) map[string]interface{} {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	user := &models.User{}

	if GetUser(username) == nil {
		return map[string]interface{}{"message": "User not found"}
	} else {
		user = GetUser(username)
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Password incorrect"}
	}

	// patients:= &models.Patient{}
	// patients = GetPatient(user.ID)

	responseUser := &models.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))

	var response = map[string]interface{}{"message": "User login successful"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}
