package handlers

import (
	"api_assessment/datab"
	"api_assessment/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// func HandleUserLogin(c *gin.Context) {
// 	body, err := ioutil.ReadAll(c.Request.Body)
// 	if err == nil {

// 	}

// 	var formattedBody models.Login
// 	err = json.Unmarshal(body, &formattedBody)

// 	login := datab.UserLogin(formattedBody.Username, formattedBody.Password)

// 	if login["message"] == "User login successful" {
// 		resp := login
// 		c.IndentedJSON(http.StatusCreated, resp)
// 	} else {

// 	}
// }

//Prototype Register Function
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := models.NewBadRequestError("Invalid JSON Body")
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := datab.CreateUser(user)
	if saveErr != nil {

		return
	}

	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := models.NewBadRequestError("Invalid JSON")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := datab.TestGetUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	if result == nil {
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(1), //Change 1 later with int(result.ID)
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString([]byte("randomtest"))
	if err != nil {
		err := models.NewInternalServerError("Login Failed")
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, result)
}
