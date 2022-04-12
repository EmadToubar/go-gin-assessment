package main

import (
	"api_assessment/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//Doctor Routes
	router.GET("/doctors", handlers.HandleGetDoctors)    //Route to view all doctors
	router.GET("/doctors/:id", handlers.HandleGetDoctor) //Route to view a specific doctor by their ID
	router.POST("/doctors", handlers.HandleAddDoctors)   //Route to add a doctor
	//router.GET("doctors/overtime", handlers.HandleGetOvertime) //Route to get Doctors with 6+ hours of appointments

	//Patient Routes
	router.GET("/patients", handlers.HandleGetPatients)    //Route to view all patients
	router.GET("/patients/:id", handlers.HandleGetPatient) //Route to get a specific patient by their ID
	router.POST("/patients", handlers.HandleAddPatients)   //Route to add a patient

	//Appointment Routes
	router.POST("/appointments/book", handlers.HandleAddAppointments)              //Route to book an appointment
	router.GET("/appointments/view/all", handlers.HandleGetAppointments)           //Route to view all appointments
	router.GET("/appointments/view/:id", handlers.HandleGetAppointment)            //Route to view a specific appointment by its ID
	router.GET("/appointments/view/patient/:id", handlers.HandleGetPatientHistory) //Route to view a specific patient's appointment history
	router.GET("/appointments/view/max", handlers.HandleGetMaxAppointments)        //Route to view the doctor with the most appointments

	//User Routes
	router.POST("/login/patient", handlers.Login)       //Route to login a Patient
	router.POST("/register", handlers.Register)         //Route to register a User
	router.POST("/testuseradd", handlers.HandleAddUser) //Route to add user (TEST ROUTE)
	//router.PATCH("/book/:id/:slot")

	router.Run("localhost:8080") //Route to run the server on port 8080
}
