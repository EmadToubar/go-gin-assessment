package models

type Doctor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	//Availability [][]string `json:"availability"`
	//Slots []Slot
}

type Slot struct {
	StartTime string
	EndTime   string
	Duration  string
}
