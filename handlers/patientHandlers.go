package handlers

import (
	"api_assessment/datab"
	"api_assessment/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetPatients(c *gin.Context) {
	patients := datab.GetPatients()

	if patients == nil || len(patients) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, patients)
	}
}

func HandleGetPatient(c *gin.Context) {
	patid := c.Param("id")

	patient := datab.GetPatient(patid)

	log.Println(patient) //Testing function REMOVE AT THE END

	if patient == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, patient)
	}
}

func HandleAddPatients(c *gin.Context) {
	var pat models.Patient

	if err := c.BindJSON(&pat); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		datab.AddPatients(pat)
		c.IndentedJSON(http.StatusCreated, pat)
	}
}
