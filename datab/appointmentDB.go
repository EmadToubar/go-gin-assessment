package datab

import (
	"api_assessment/models"
	"fmt"
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

	if appointment.Duration > 120 {
		log.Println("Exceeded booking duration")
		panic("Exceeded booking duration")

	} else {
		//Insert rest of code here when everythign works
	}

	insert, err := db.Queryx(
		"INSERT INTO appointment (time, id, doctorid, patientid, duration, timestart, timeend) VALUES (($1),($2),($3), ($4), ($5), ($6), ($7))",
		appointment.Time, appointment.ID, appointment.DoctorID, appointment.PatientID, appointment.Duration, appointment.TimeStart, appointment.TimeEnd)

	// if there is an error inserting, handle it
	if err == nil {
		//panic(err.Error())
		return nil
	} else {

	}
	defer insert.Close()
	return &res
}

//Function to list all the doctors in the DB
func GetAppointments() []models.Appointment {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	results, err := db.Queryx("SELECT * FROM appointment")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	appoint := []models.Appointment{}
	for results.Next() {
		var a models.Appointment
		// for each row, scan into the Appointment struct
		err = results.Scan(&a.Time, &a.ID, &a.DoctorID, &a.PatientID, &a.Duration, &a.TimeStart, &a.TimeEnd)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the doctor into doctors array
		appoint = append(appoint, a)
	}

	return appoint

}

//Function to get an appointment by its ID in the DB
func GetAppointment(appointid int) *models.Appointment {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	appoint := &models.Appointment{}

	results, err := db.Queryx("SELECT * FROM appointment where id=($1)", appointid)
	log.Println(results) //Test function REMOVE AT THE END

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&appoint.Time, &appoint.ID, &appoint.DoctorID, &appoint.PatientID, &appoint.Duration, &appoint.TimeStart, &appoint.TimeEnd)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return appoint
}

//Function to get a patient's appointment history
func GetPatientHistory(patientid string) []models.Appointment {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	results, err := db.Queryx("SELECT * FROM appointment where patientid=($1)", patientid)
	log.Println(results) //Test function REMOVE AT THE END

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	appointarr := []models.Appointment{}
	for results.Next() {
		var a models.Appointment
		// for each row, scan into the Appointment struct
		err = results.Scan(&a.Time, &a.ID, &a.DoctorID, &a.PatientID, &a.Duration, &a.TimeStart, &a.TimeEnd)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the doctor into appointment array
		appointarr = append(appointarr, a)
	}

	return appointarr
}

//Function to get the doctor with the most appointments
func GetMaxAppointments() []models.CountResponse {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	results, err := db.Queryx("SELECT doctorid, COUNT(doctorid) FROM appointment GROUP BY doctorid ORDER BY COUNT(doctorid) DESC LIMIT 1 ")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	response := []models.CountResponse{}
	for results.Next() {
		var a models.CountResponse
		// for each row, scan into the Appointment struct
		err = results.Scan(&a.DoctorId, &a.Count)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the doctor into doctors array
		response = append(response, a)
	}

	return response

}
