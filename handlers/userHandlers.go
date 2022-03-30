package handlers

// import (
// 	"api_assessment/datab"
// 	"api_assessment/models"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

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
