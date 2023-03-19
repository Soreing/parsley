package models

import "github.com/Soreing/fastjson/tests/externals"

type Employee struct {
	Id          string             `json:"id"`
	Person      Person             `json:"person"`
	Devices     []externals.Device `json:"devices"`
	IsActive    bool               `json:"isActive"`
	Rating      float64            `json:"rating"`
	LineManager *Employee          `json:"lineManager"`
	Tags        []string           `json:"tags"`
}

type EmployeeList []Employee
