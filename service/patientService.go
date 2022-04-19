package service

import (
	"api_assessment/datab"
	"api_assessment/models"
)

type PatientService interface {
	GetPatients() []models.Patient
	GetPatient(patid string) *models.Patient
	AddPatients(patient models.Patient)
}

type patientService struct {
	patientDb datab.PatientDB
}

func (pt *patientService) GetPatients() []models.Patient {
	return pt.GetPatients()
}

func (pt *patientService) GetPatient(patid string) *models.Patient {
	return pt.GetPatient(patid)
}

func (pt *patientService) AddPatients(patient models.Patient) {
	pt.AddPatients(patient)
}

func PatientServiceProvider(db *datab.PatientDB) PatientService {
	return &patientService{
		patientDb: *db,
	}
}
