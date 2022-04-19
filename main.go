package main

import (
	"api_assessment/datab"
	"api_assessment/handlers"
	"api_assessment/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	ctx := datab.DbCtxProvider()
	defer ctx.Close()
	//################### User Flow #######################
	userDB := datab.UserDBProvider(ctx)
	userService := service.NewUserService(userDB)
	userHandlers := handlers.NewUserHandler(userService)
	baseUrl := router.Group("")
	userHandlers.SetupRoutes(baseUrl)
	//################### Doctor Flow #######################
	docDB := datab.DoctorDBProvider(ctx)
	docService := service.DoctorServiceProvider(&docDB)
	docHandler := handlers.DoctorHandlerProvider(docService)
	docHandler.SetupRoutes(baseUrl)
	//################### Doctor Flow #######################
	appDB := datab.AppointmentDBProvider(ctx)
	appService := service.AppointmentServiceProvider(&appDB)
	appHandler := handlers.AppointmentHandlerProvider(appService)
	appHandler.SetupRoutes(baseUrl)
	//################### Patient Flow #######################
	patDB := datab.PatientDBProvider(ctx)
	patService := service.PatientServiceProvider(&patDB)
	patHandler := handlers.PatientHandlerProvider(patService)
	patHandler.SetupRoutes(baseUrl)
	//#####################################################
	//Doctor Routes
	// router.GET("/doctors", handlers.HandleGetDoctors)    //Route to view all doctors
	// router.GET("/doctors/:id", handlers.HandleGetDoctor) //Route to view a specific doctor by their ID
	// router.POST("/doctors", handlers.HandleAddDoctors)   //Route to add a doctor
	//router.GET("doctors/overtime", handlers.HandleGetOvertime) //Route to get Doctors with 6+ hours of appointments

	//Patient Routes
	// router.GET("/patients", handlers.HandleGetPatients)    //Route to view all patients
	// router.GET("/patients/:id", handlers.HandleGetPatient) //Route to get a specific patient by their ID
	// router.POST("/patients", handlers.HandleAddPatients)   //Route to add a patient

	//Appointment Routes
	// router.POST("/appointments/book", handlers.HandleAddAppointments)              //Route to book an appointment
	// router.GET("/appointments/view/all", handlers.HandleGetAppointments)           //Route to view all appointments
	// router.GET("/appointments/view/:id", handlers.HandleGetAppointment)            //Route to view a specific appointment by its ID
	// router.GET("/appointments/view/patient/:id", handlers.HandleGetPatientHistory) //Route to view a specific patient's appointment history
	// router.GET("/appointments/view/max", handlers.HandleGetMaxAppointments)        //Route to view the doctor with the most appointments

	//User Routes
	// router.POST("/login/patient", handlers.Login) //Route to login a Patient
	// router.POST("/register", handlers.Register)   //Route to register a User
	// router.POST("/testuseradd", handlers.HandleAddUser) //Route to add user (TEST ROUTE)
	//router.PATCH("/book/:id/:slot")

	router.Run("localhost:8080") //Route to run the server on port 8080
}
