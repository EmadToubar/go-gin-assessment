package handlers

import (
	"api_assessment/datab"
	"api_assessment/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Prototype Register Function
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := models.NewBadRequestError("Invalid JSON Body")
		c.JSON(err.Status, err)
		return
	}

	log.Println("Testing user binding in Register handler:", user)

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

//Test Function
func HandleAddUser(c *gin.Context) {
	var use models.User

	if err := c.BindJSON(&use); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		datab.AddUsers(use)

		c.IndentedJSON(http.StatusCreated, use)
	}
}
