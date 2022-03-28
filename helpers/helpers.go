package helpers

import (
	"api_assessment/models"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}

	return string(hashed)
}

func Validation(values []models.Validation) bool {
	username := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z0-9]$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}

		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}

		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}
