package service

import "api_assessment/datab"

type AppointmentService interface{}

type appointmentService struct {
	appointmentDb datab.AppointmentDB
}

func AppointmentServiceProvider(db *datab.AppointmentDB) AppointmentService {
	return &appointmentService{
		appointmentDb: *db,
	}
}
