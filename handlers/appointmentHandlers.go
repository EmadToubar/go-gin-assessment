package handlers

import (
	"api_assessment/datab"
	"api_assessment/models"
	"api_assessment/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler interface {
	HandleGetAppointments(c *gin.Context)
	HandleGetAppointment(c *gin.Context)
	HandleGetPatientHistory(c *gin.Context)
	HandleGetMaxAppointments(c *gin.Context)
	SetupRoutes(r *gin.RouterGroup)
}

type appointmentHandler struct {
	service service.AppointmentService
}

func AppointmentHandlerProvider(service service.AppointmentService) AppointmentHandler {
	return &appointmentHandler{
		service: service,
	}
}

func (ah *appointmentHandler) SetupRoutes(r *gin.RouterGroup) {
	r.GET("/appointments", ah.HandleGetAppointments)
	r.GET("/appointments/:id", ah.HandleGetAppointment)
	r.GET("/appointments/:id/history", ah.HandleGetPatientHistory)
	r.GET("/appointments/max", ah.HandleGetMaxAppointments)
}

func (ah *appointmentHandler) HandleAddAppointments(c *gin.Context) {
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

func (ah *appointmentHandler) HandleGetAppointments(c *gin.Context) {
	appoint := datab.GetAppointments()

	if appoint == nil || len(appoint) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, appoint)
	}
}

func (ah *appointmentHandler) HandleGetAppointment(c *gin.Context) {
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

func (ah *appointmentHandler) HandleGetPatientHistory(c *gin.Context) {
	patid := c.Param("id")

	a := datab.GetPatientHistory(patid)

	log.Println(a) //Testing function REMOVE AT THE END

	if a == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, a)
	}
}

func (ah *appointmentHandler) HandleGetMaxAppointments(c *gin.Context) {
	appoint := datab.GetMaxAppointments()

	if appoint == nil || len(appoint) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, appoint)
	}
}
