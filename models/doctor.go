package models

type Doctor struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Role         string   `json:"role"`
	Availability []string `json:"availability"`
	Patients     []Patient
}
