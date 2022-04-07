package datab

import (
	"api_assessment/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Function to list all the doctors in the DB
func GetDoctors() []models.Doctor {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	results, err := db.Queryx("SELECT * FROM doctor")

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
func GetDoctor(docid string) *models.Doctor {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	doc := &models.Doctor{}

	results, err := db.Queryx("SELECT * FROM doctor where id=($1)", docid)
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
func AddDoctors(doctor models.Doctor) {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	defaultSlots := [...]string{
		"8:00",
		"8:15",
		"8:30",
		"8:45",
		"9:00"}

	insert, err := db.Queryx(
		"INSERT INTO doctor (id, name, role, availability) VALUES (($1),($2),($3), ($4))",
		doctor.ID, doctor.Name, doctor.Role, defaultSlots)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func CheckAvailability(doctors models.Doctor, slot string) bool {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()

	for i := 0; i < len(doctors.Availability); i++ {
		if doctors.Availability[i] == slot {
			log.Println("Slot is taken") //Placeholder code
			return false
		}
	}
	log.Println("Slot is free") //Placeholder code
	return true
}

func removeSlot(r []string, s string) []string {
	for i := 0; i < len(r); i++ {
		if r[i] == s {
			copy(r[i:], r[i+1:])
			r[len(r)-1] = ""
			r = r[:len(r)-1]
		}
	}
	return r
}

func BookSlot(doctors models.Doctor, slot string) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	defer db.Close()
	if CheckAvailability(doctors, slot) == true {
		doctors.Availability = removeSlot(doctors.Availability, slot)
	} else {
		log.Println("Slot already booked.") //Placeholder code
	}

}
