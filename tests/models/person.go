package models

import "time"

type Person struct {
	Fname string    `json:"fname"`
	Lname string    `json:"lname"`
	DOB   time.Time `json:"dob"`
}
