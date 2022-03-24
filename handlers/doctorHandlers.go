package handlers

import (
	"api_assessment/datab"
	"api_assessment/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetDoctors(c *gin.Context) {
	doctors := datab.GetDoctors()

	if doctors == nil || len(doctors) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, doctors)
	}
}

func HandleGetDoctor(c *gin.Context) {
	docid := c.Param("id")

	doctor := datab.GetDoctor(docid)

	log.Println(doctor) //Testing function REMOVE AT THE END

	if doctor == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, doctor)
	}
}

func HandleAddDoctors(c *gin.Context) {
	var doc models.Doctor

	if err := c.BindJSON(&doc); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		datab.AddDoctors(doc)
		c.IndentedJSON(http.StatusCreated, doc)
	}
}
