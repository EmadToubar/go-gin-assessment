package handlers

import (
	"api_assessment/service"

	"github.com/gin-gonic/gin"
)

type PatientHandler interface {
	HandleGetPatients(c *gin.Context)
	HandleGetPatient(c *gin.Context)
	HandleAddPatients(c *gin.Context)
	SetupRoutes(r *gin.RouterGroup)
}

type patientHandler struct {
	service service.PatientService
}

func (ph *patientHandler) HandleGetPatients(c *gin.Context) {
	// patients := datab.GetPatients()

	// if patients == nil || len(patients) == 0 {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// } else {
	// 	c.IndentedJSON(http.StatusOK, patients)
	// }
}

func (ph *patientHandler) HandleGetPatient(c *gin.Context) {
	// patid := c.Param("id")

	// patient := datab.GetPatient(patid)

	// log.Println(patient) //Testing function REMOVE AT THE END

	// if patient == nil {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// } else {
	// 	c.IndentedJSON(http.StatusOK, patient)
	// }
}

func (ph *patientHandler) HandleAddPatients(c *gin.Context) {
	// var pat models.Patient

	// if err := c.BindJSON(&pat); err != nil {
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// } else {
	// 	datab.AddPatients(pat)
	// 	c.IndentedJSON(http.StatusCreated, pat)
	// }
}

func (ph *patientHandler) SetupRoutes(r *gin.RouterGroup) {
	r.GET("/patients", ph.HandleGetPatients)
	r.GET("/patients/:id", ph.HandleGetPatient)
	r.POST("/patients", ph.HandleAddPatients)
}

func PatientHandlerProvider(service service.PatientService) PatientHandler {
	return &patientHandler{
		service: service,
	}
}
