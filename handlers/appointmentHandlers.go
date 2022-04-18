package handlers

import (
	"api_assessment/models"
	"api_assessment/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler interface {
	HandleAddAppointments(c *gin.Context)
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
	r.POST("/appointments/book", ah.HandleAddAppointments)
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
		result, err := ah.service.AddAppointments(appoint)
		if result == nil {

		}
		if err != nil {
		}
		c.IndentedJSON(http.StatusCreated, appoint)
	}
}

func (ah *appointmentHandler) HandleGetAppointments(c *gin.Context) {
	appoint, err := ah.service.GetAppointments()

	if appoint == nil || len(appoint) == 0 || err != nil {
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

	a, err := ah.service.GetAppointment(appointidint)

	log.Println(a) //Testing function REMOVE AT THE END

	if a == nil || err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, a)
	}
}

func (ah *appointmentHandler) HandleGetPatientHistory(c *gin.Context) {
	patid := c.Param("id")

	a, err := ah.service.GetPatientHistory(patid)

	log.Println(a) //Testing function REMOVE AT THE END

	if a == nil || err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, a)
	}
}

func (ah *appointmentHandler) HandleGetMaxAppointments(c *gin.Context) {
	appoint, err := ah.service.GetMaxAppointments()

	if appoint == nil || len(appoint) == 0 || err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, appoint)
	}
}
