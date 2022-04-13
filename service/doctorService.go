package service

import (
	"api_assessment/datab"
)

type DoctorService interface{}

type doctorService struct {
	doctorDb datab.DoctorDB
}

func DoctorServiceProvider(db *datab.DoctorDB) DoctorService {
	return &doctorService{
		doctorDb: *db,
	}
}
