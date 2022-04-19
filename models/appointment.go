package models

type Appointment struct {
	Time       string  `json:"time"`
	SchDoctor  Doctor  `json:"doctor,omitempty"`
	SchPatient Patient `json:"patient,omitempty"`
	ID         int     `json:"id"`
	DoctorID   string  `json:"doctorid"`
	PatientID  string  `json:"patientid"`
	Duration   int     `json:"duration"`
	TimeStart  string  `json:"timestart"`
	TimeEnd    string  `json:"timeend"`
}

type CountResponse struct {
	DoctorId string
	Count    int
}
