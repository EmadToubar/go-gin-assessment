package service

import "api_assessment/datab"

type PatientService interface{}

type patientService struct {
	patientDb datab.PatientDB
}

func PatientServiceProvider(db *datab.PatientDB) PatientService {
	return &patientService{
		patientDb: *db,
	}
}
