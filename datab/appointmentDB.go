package datab

import (
	"api_assessment/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Function to add an appointment to the DB
func AddAppointments(appointment models.Appointment) *models.Appointment {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	res := appointment
	res.SchPatient = *GetPatient(appointment.PatientID)
	res.SchDoctor = *GetDoctor(appointment.DoctorID)
	log.Println(res) //Testing to see if fetching is working

	insert, err := db.Queryx(
		"INSERT INTO appointment (id, doctorid, patientid, duration, timestart, timeend) VALUES (($1),($2),($3), ($4), ($5), ($6))",
		appointment.ID, appointment.DoctorID, appointment.PatientID, appointment.Duration, appointment.TimeStart, appointment.TimeEnd)

	// if there is an error inserting, handle it
	if err == nil {
		//panic(err.Error())
		return nil
	} else {

	}
	defer insert.Close()
	return &res
}
