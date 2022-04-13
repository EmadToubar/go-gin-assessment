package service

import (
	"api_assessment/datab"
	"api_assessment/models"
)

type UserService interface {
	RegisterUser(models.User) (*models.User, error)
	Login(models.User) (*models.User, *string, error)
	AddUser(*models.User)
}

type userService struct {
	userDb datab.UserDB
}

func (us *userService) RegisterUser(user models.User) (*models.User, error) {
	//TODO: call the appropriate function in the datab package from the userDB interface
	return nil, nil
}

func (us *userService) Login(user models.User) (*models.User, *string, error) { //return the user model, jwt token, error if any
	//TODO: call the appropriate function in the datab package from the userDB interface
	// FEEEDBACK: all od this can be abstracted in the service layer where you would ideally want to handle all of the logic
	// result, getErr := datab.TestGetUser(user)
	// if getErr != nil {
	// 	c.JSON(getErr.Status, getErr)
	// 	return
	// }
	// if result == nil {
	// }
	// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer:    result.ID, //Change 1 later with int(result.ID)
	// 	ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	// })

	// token, err := claims.SignedString([]byte("randomtest"))
	// if err != nil {
	// 	err := models.NewInternalServerError("Login Failed")
	// 	c.JSON(err.Status, err)
	// 	return
	// }

	// c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	// c.JSON(http.StatusOK, result)
	return nil, nil, nil
}

func (us *userService) AddUser(user *models.User) {
}

func NewUserService(userDb datab.UserDB) UserService {
	return &userService{
		userDb: userDb,
	}
}
