package service

import (
	"api_assessment/datab"
	"api_assessment/models"
)

type AppointmentService interface {
	AddAppointments(models.Appointment) (*models.Appointment, error)
	GetAppointments() ([]models.Appointment, error)
	GetAppointment(int) (*models.Appointment, error)
	GetPatientHistory(string) ([]models.Appointment, error)
	GetMaxAppointments() ([]models.CountResponse, error)
}

type appointmentService struct {
	appointmentDb datab.AppointmentDB
}

func (ap *appointmentService) AddAppointments(appoint models.Appointment) (*models.Appointment, error) {
	// result, err := ap.AddAppointments(appoint)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return ap.AddAppointments(appoint)

}

func (ap *appointmentService) GetAppointments() ([]models.Appointment, error) {
	// result, err := ap.GetAppointments()
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return ap.GetAppointments()
}

func (ap *appointmentService) GetAppointment(appointid int) (*models.Appointment, error) {
	// result, err := ap.GetAppointment(appointid)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return ap.GetAppointment(appointid)

}

func (ap *appointmentService) GetPatientHistory(patientid string) ([]models.Appointment, error) {
	// result, err := ap.GetPatientHistory(patientid)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return ap.GetPatientHistory(patientid)
}

func (ap *appointmentService) GetMaxAppointments() ([]models.CountResponse, error) {
	// result, err := ap.GetMaxAppointments()
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return ap.GetMaxAppointments()
}

func AppointmentServiceProvider(db *datab.AppointmentDB) AppointmentService {
	return &appointmentService{
		appointmentDb: *db,
	}
}
