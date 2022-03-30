package handlers

import (
	"api_assessment/datab"
	"api_assessment/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAddAppointments(c *gin.Context) {
	var appoint models.Appointment

	if err := c.BindJSON(&appoint); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		test := datab.AddAppointments(appoint)
		log.Println(test)
		if test == test {
		}
		c.IndentedJSON(http.StatusCreated, appoint)
	}
}
