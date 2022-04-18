package datab

import (
	"api_assessment/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PatientDB interface {
	GetPatients() []models.Patient
	GetPatient(patid string) *models.Patient
	AddPatients(patient models.Patient)
}

type patientDB struct {
	db *sqlx.DB
}

func PatientDBProvider(db *sqlx.DB) PatientDB {
	return &patientDB{
		db: db,
	}
}

//Function to list all the patients in the DB
func (pd *patientDB) GetPatients() []models.Patient {

	results, err := pd.db.Queryx("SELECT * FROM patient")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	patients := []models.Patient{}
	for results.Next() {
		var pat models.Patient
		// for each row, scan into the Patient struct
		err = results.Scan(&pat.ID, &pat.Name, &pat.Role)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the patient into patients array
		patients = append(patients, pat)
	}

	return patients

}

//Function to get a patient by their ID from the DB
func (pd *patientDB) GetPatient(patid string) *models.Patient {

	pat := &models.Patient{}

	results, err := pd.db.Queryx("SELECT * FROM patient where id=($1)", patid)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&pat.ID, &pat.Name, &pat.Role)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return pat
}

//Function to add a patient to the DB
func (pd *patientDB) AddPatients(patient models.Patient) {

	insert, err := pd.db.Queryx(
		"INSERT INTO patient (id, name, role) VALUES (($1),($2),($3))",
		patient.ID, patient.Name, patient.Role)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
