package main

import (
	"api_assessment/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/doctors", handlers.HandleGetDoctors)                 //Route to view all doctors
	router.GET("/doctors/:id", handlers.HandleGetDoctor)              //Route to view a specific doctor by their ID
	router.POST("/doctors", handlers.HandleAddDoctors)                //Route to add a doctor
	router.GET("/patients", handlers.HandleGetPatients)               //Route to view all patients
	router.GET("/patients/:id", handlers.HandleGetPatient)            //Route to get a specific patient by their ID
	router.POST("/patients", handlers.HandleAddPatients)              //Route to add a patient
	router.POST("/appointments/book", handlers.HandleAddAppointments) //Route to book an appointment
	router.Run("localhost:8080")                                      //Route to run the server on port 8080
}