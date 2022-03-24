package models

type Appointment struct {
	Time string `json:"time"`
	// schDoctor  Doctor
	// schPatient Patient
	ID        int    `json:"id"`
	DoctorID  string `json:"doctorid"`
	PatientID string `json:"patientid"`
	Duration  string `json:"duration"`
	TimeStart string `json:"timestart"`
	TimeEnd   string `json:"timeend"`
}
