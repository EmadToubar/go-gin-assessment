package datab

import (
	"api_assessment/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type doctorDB struct {
	db *sqlx.DB
}

type DoctorDB interface {
	GetDoctors() []models.Doctor
	GetDoctor(docid string) *models.Doctor
	AddDoctors(doctor models.Doctor) error
	CheckAvailability(doctors models.Doctor, appoint models.Appointment, slot string, docid string) []models.Slot
	BookSlot(doctors models.Doctor, slot string) error
}

func DoctorDBProvider(ctx *sqlx.DB) DoctorDB {
	return &doctorDB{
		db: ctx,
	}
}

//Function to list all the doctors in the DB
func (dd *doctorDB) GetDoctors() []models.Doctor {
	results, err := dd.db.Queryx("SELECT * FROM doctor")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	doctors := []models.Doctor{}
	for results.Next() {
		var doc models.Doctor
		// for each row, scan into the Doctor struct
		err = results.Scan(&doc.ID, &doc.Name, &doc.Role)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the doctor into doctors array
		doctors = append(doctors, doc)
	}

	return doctors

}

//Function to get a doctor by their ID in the DB
func (dd *doctorDB) GetDoctor(docid string) *models.Doctor {
	doc := &models.Doctor{}
	results, err := dd.db.Queryx("SELECT * FROM doctor where id=($1)", docid)
	log.Println(results) //Test function REMOVE AT THE END

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&doc.ID, &doc.Name, &doc.Role)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return doc
}

//Function to add a doctor to the DB
func (dd *doctorDB) AddDoctors(doctor models.Doctor) error {
	if insert, err := dd.db.Queryx(
		"INSERT INTO doctor (id, name, role) VALUES (($1),($2),($3))",
		doctor.ID, doctor.Name, doctor.Role); err != nil {
		return err
	} else {
		return insert.Close()
	}
}

func (dd *doctorDB) CheckAvailability(doctors models.Doctor, appoint models.Appointment, slot string, docid string) []models.Slot {
	//Emad To-Do List:
	//1. Change slots from strings to other form to allow for calculations

	//Create a temp object with start time of beginning of day and end time of start time of first eppointment
	//Create objects where

	// availableSlots := []slots{}]
	//----------First Psuedo code I worked on (IGNORE)-----------
	// availableSlots := [...]models.Slot{}
	// bookedSlotsString := doctors.Availability
	// bookedSlotsTime := [...]time.Time{}
	// var err error
	// for i := 0; i <= len(bookedSlotsString); i++ {
	// 	bookedSlotsTime[i], err = time.Parse("150405", bookedSlotsString[i])
	// }
	// if err != nil {

	// }
	// var j int
	// for i := 0; i <= len(doctors.Availability); i++ {
	// 	j = 0
	// 	if doctors.Availability[0] == "8:00" && doctors.Availability[0][0] == "free" {
	// 		availableSlots[i].StartTime = "8:00"
	// 		while(doctors.Availability[i+1][0] != "booked")
	// 		{
	// 			j++
	// 		}
	// 		availableSlots[i].EndTime = doctors.Availability[j][0]
	// 	} else {

	// 	}

	// }
	//-------------------------------------------------------------------
	//-----------------Second Code------------------------------------//
	//To do:
	//1. Apply appointment DB function to this function
	//2. Store all appointments related to doctor in doctorAppointments
	//3. Create a for loop that goes through the appointments and sets the start time of the availableSlots as the end time of the appointment and the end time of the avaialbleSlots as the start time of the next appointment
	//4. Return the availableSlots object

	//Todo (Karim's List): These are the steps to check if the slot is available
	//1. Get All of the doctors  booked slots for the day
	//2. calculate the difference between the doctor's first slot and the start of the day i.e. 8am and add it the available slots array
	//3. loop through the booked slots array and calculate the difference between the end time of slot i and start time of slot (i+1)
	//4. when you reach the last item in the booked slots array, calculate the difference between the end time of slot i and the end of the day i.e. 5pm and add it to the available slots array
	//5. return the available slots array as this will have a list of all the available slots for the day

	doctorAppointments := dd.GetDocAppointment(docid)
	availableSlots := [...]models.Slot{}
	for i := 0; i <= len(doctorAppointments); i++ {
		if doctorAppointments[0].TimeStart == "8:00" && i == 0 {
			availableSlots[0].StartTime = doctorsAppointments[0].TimeEnd
			availableSlots[0].EndTime = doctorsAppointments[1].TimeStart
		} else if doctorAppointments[0].TimeStart == "8:00" != i == 0 {
			availableSlots[0].StartTime = "8:00"
			availableSlots[0].EndTime = doctorsAppointments[0].TimeStart
		} else {
			availableSlots[i].StartTime = doctorAppointments[i].TimeEnd
			availableSlots[i].EndTime = doctorAppointments[i+1].TimeStart
		}

	}

	log.Println("Slot is free") //Placeholder code
	return availableSlots[:]
}

func (dd *doctorDB) removeSlot(r []string, s string) []string {
	for i := 0; i < len(r); i++ {
		if r[i] == s {
			copy(r[i:], r[i+1:])
			r[len(r)-1] = ""
			r = r[:len(r)-1]
		}
	}
	return r
}

func (dd *doctorDB) BookSlot(doctors models.Doctor, slot string) error {
	//Todo: These are the steps to book a slot
	//1. Query the database to get to see if the doctor has any slots booked between the start and end of the slot
	//2. If there are no slots booked, Check the duration of the slot and if between 15 mins and 120 mins book it.
	//3. if any of these checks fail, return an error
	return nil
}
