package datab

import (
	"api_assessment/helpers"
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

//Function for creating user account
func CreateAccount() {
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

	}

}

// //Function for logging in the user
// func UserLogin(username string, pass string) map[string]interface{} {
// 	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
// 	if err != nil {
// 		log.Fatalln(err)
// 	} //Connecting to database

// 	db.MustExec(schema)

// 	defer db.Close()

// 	user := &models.User{}

// 	if GetUser(username) == nil {
// 		return map[string]interface{}{"message": "User not found"}
// 	} else {
// 		user = GetUser(username)
// 	}

// 	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

// 	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
// 		return map[string]interface{}{"message": "Password incorrect"}
// 	}

// 	// patients:= &models.Patient{}
// 	// patients = GetPatient(user.ID)

// 	responseUser := &models.User{
// 		ID:       user.ID,
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		Password: user.Password,
// 		Role:     user.Role,
// 	}

// 	tokenContent := jwt.MapClaims{
// 		"user_id": user.ID,
// 		"role":    user.Role,
// 		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
// 	}
// 	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
// 	token, err := jwtToken.SignedString([]byte("TokenPassword"))

// 	var response = map[string]interface{}{"message": "User login successful"}
// 	response["jwt"] = token
// 	response["data"] = responseUser

// 	return response
// }

// func UserRegister(username string, email string, pass string) map[string]interface{} {
// 	valid := helpers.Validation(
// 		[]models.Validation{
// 			{Value: username, Valid: "username"},
// 			{Value: email, Valid: "email"},
// 			{Value: pass, Valid: "password"},
// 		})

// 	if valid {

// 		db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
// 		if err != nil {
// 			log.Fatalln(err)
// 		} //Connecting to database

// 		db.MustExec(schema)

// 		defer db.Close()

// 		generatedPassword := helpers.HashAndSalt([]byte(pass))
// 		user := models.User{Name: username, Email: email, Password: generatedPassword}

// 		patient := models.Patient{ID: "50", Name: username, Role: "PATIENT"}

// 		AddPatients(patient)
// 		AddUsers(user)

// 		//patients:= []models.Patient{}

// 		return map[string]interface{}{"message": "User registered."}

// 	} else {
// 		return map[string]interface{}{"message": "Invalid values."}
// 	}

// }

func CreateUser(user models.User) (*models.User, *models.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, models.NewBadRequestError("Failed to Encrypt Password")
	}

	user.Password = string(pwSlice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func TestGetUser(user models.User) (*models.User, *models.RestErr) {
	result := &models.User{Email: user.Email}

	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, models.NewBadRequestError("Failed to Decrypt Password")
	}

	resultWp := &models.User{ID: result.ID, Name: result.Name, Email: result.Email, Role: result.Role}
	return resultWp, nil
}

func (user *models.User) Validate() *models.RestErr {
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

func (user *models.User) Save() *models.RestErr {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(*GiveSchema())

	defer db.Close()

	insert, err := db.Queryx(
		"INSERT INTO users (id, name, email, password, role) VALUES (($1),($2),($3), ($4), ($5))",
		user.ID, user.Name, user.Email, user.Password, "PATIENT")

	if insert == nil {

	}

	return nil

}

func (user *models.User) GetByEmail() *models.RestErr {
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

}
