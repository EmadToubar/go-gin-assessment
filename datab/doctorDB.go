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
	CheckAvailability(doctors models.Doctor, slot string) bool
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

func (dd *doctorDB) CheckAvailability(doctors models.Doctor, slot string) bool {
	// availableSlots := []slots{}
	//Todo: These are the steps to check if the slot is available
	//1. Get All of the doctors  booked slots for the day
	//2. calculate the difference between the doctor's first slot and the start of the day i.e. 8am and add it the available slots array
	//3. loop through the booked slots array and calculate the difference between the end time of slot i and start time of slot (i+1)
	//4. when you reach the last item in the booked slots array, calculate the difference between the end time of slot i and the end of the day i.e. 5pm and add it to the available slots array
	//5. return the available slots array as this will have a list of all the available slots for the day
	log.Println("Slot is free") //Placeholder code
	return true
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
