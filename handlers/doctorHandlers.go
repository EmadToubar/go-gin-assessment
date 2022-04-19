package handlers

import (
	"api_assessment/service"

	"github.com/gin-gonic/gin"
)

type DoctorHandler interface {
	SetupRoutes(*gin.RouterGroup)
	HandleGetDoctors(c *gin.Context)
	HandleGetDoctor(c *gin.Context)
	HandleAddDoctors(c *gin.Context)
}
type doctorHandler struct {
	service service.DoctorService
}

func DoctorHandlerProvider(service service.DoctorService) DoctorHandler {
	return &doctorHandler{
		service: service,
	}
}

func (dh *doctorHandler) HandleGetDoctors(c *gin.Context) {
	// doctors := datab.GetDoctors()

	// if doctors == nil || len(doctors) == 0 {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// } else {
	// 	c.IndentedJSON(http.StatusOK, doctors)
	// }
}

func (dh *doctorHandler) HandleGetDoctor(c *gin.Context) {
	// docid := c.Param("id")

	// doctor := datab.GetDoctor(docid)

	// log.Println(doctor) //Testing function REMOVE AT THE END

	// if doctor == nil {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// } else {
	// 	c.IndentedJSON(http.StatusOK, doctor)
	// }
}

func (dh *doctorHandler) HandleAddDoctors(c *gin.Context) {
	// var doc models.Doctor

	// if err := c.BindJSON(&doc); err != nil {
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// } else {
	// 	datab.AddDoctors(doc)
	// 	c.IndentedJSON(http.StatusCreated, doc)
	// }
}

//Test Function for Slots
func (dh *doctorHandler) HandleBookSlot(c *gin.Context) {
	// slot := c.Param("slot")
	// docid := c.Param("id")

	// var doc models.Doctor

	// doc = *datab.GetDoctor(docid)
	// datab.BookSlot(doc, slot)
}

func (dh *doctorHandler) SetupRoutes(c *gin.RouterGroup) {
	c.GET("/doctors", dh.HandleGetDoctors)
	c.GET("/doctors/:id", dh.HandleGetDoctor)
	c.POST("/doctors", dh.HandleAddDoctors)
	// c.GET("/doctors/:id/:slot", dh.HandleBookSlot)
}
