package handlers

import (
	"api_assessment/datab"
	"api_assessment/models"
	"log"
	"net/http"
	"strconv"

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

func HandleGetAppointments(c *gin.Context) {
	appoint := datab.GetAppointments()

	if appoint == nil || len(appoint) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, appoint)
	}
}

func HandleGetAppointment(c *gin.Context) {
	appointid := c.Param("id")
	appointidint, err := strconv.Atoi(appointid)
	if err == nil {
	}

	a := datab.GetAppointment(appointidint)

	log.Println(a) //Testing function REMOVE AT THE END

	if a == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, a)
	}
}

func HandleGetPatientHistory(c *gin.Context) {
	patid := c.Param("id")

	a := datab.GetPatientHistory(patid)

	log.Println(a) //Testing function REMOVE AT THE END

	if a == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, a)
	}
}

func HandleGetMaxAppointments(c *gin.Context) {
	appoint := datab.GetMaxAppointments()

	if appoint == nil || len(appoint) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, appoint)
	}
}
