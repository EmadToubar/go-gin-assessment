package service

import (
	"api_assessment/datab"
	"api_assessment/models"
)

type DoctorService interface {
	GetDoctors() []models.Doctor
	GetDoctor(docid string) *models.Doctor
	AddDoctors(doctor models.Doctor) error
}

type doctorService struct {
	doctorDb datab.DoctorDB
}

func (dc *doctorService) GetDoctors() []models.Doctor {
	return dc.GetDoctors()
}

func (dc *doctorService) GetDoctor(docid string) *models.Doctor {
	return dc.GetDoctor(docid)
}

func (dc *doctorService) AddDoctors(doctor models.Doctor) error {
	return dc.AddDoctors(doctor)
}

func DoctorServiceProvider(db *datab.DoctorDB) DoctorService {
	return &doctorService{
		doctorDb: *db,
	}
}
