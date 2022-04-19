package handlers

import (
	"api_assessment/models"
	"api_assessment/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	SetupRoutes(c *gin.RouterGroup)
}
type userHandler struct {
	service service.UserService
}

//Prototype Register Function
func (uh *userHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		err := models.NewBadRequestError("Invalid JSON Body")
		c.JSON(err.Status, err)
	} else {
		log.Println("Testing user binding in Register handler:", user)
		if result, saveErr := uh.service.RegisterUser(user); saveErr != nil {
			//TODO: Handle the error and return it to the user
		} else {
			c.JSON(http.StatusOK, result)
		}
	}

}

func (uh *userHandler) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := models.NewBadRequestError("Invalid JSON")
		c.JSON(err.Status, err)
		return
	} else {
		log.Println("Testing user binding in Login handler:", user)
		if result, token, err := uh.service.Login(user); err != nil {
		} else {
			c.JSON(http.StatusOK, gin.H{"token": token, "details": result})
		}
	}

}

func (uh *userHandler) SetupRoutes(c *gin.RouterGroup) {
	c.POST("/register", uh.Register)
	c.POST("/login", uh.Login)
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}

//Test Function
// func (uh *userHandler) HandleAddUser(c *gin.Context) { // i dont think you needed this cause register should be your only entry point to creating an app
// 	var use models.User

// 	if err := c.BindJSON(&use); err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 	} else {
// 		uh.service.AddUser(&use)
// 		c.IndentedJSON(http.StatusCreated, use)
// 	}
// }
