package main

import (
	"testing"
	"time"

	"github.com/Soreing/parsley"
	"github.com/Soreing/parsley/tests/externals"
	"github.com/Soreing/parsley/tests/models"
)

const EmployeeJSON = `{
	"id": "8becdcce-e762-40cf-b73c-092209f70a30",
	"person": {
		"fname": "John",
		"lname": "Smith",
		"dob": "1998-04-25T00:00:00Z"
	},
	"task": {
		"name": "Data Entry",
		"assigned": "2023-03-10:10:00:00Z",
		"state": 1
	},
	"devices": [
		{
			"name": "Work Computer",
			"type": 0
		},
		{
			"name": "Work Mobile",
			"type": 1
		}
	],
	"isActive": true,
	"rating": 8.65,
	"lineManager": null,
	"tags": ["punctual", "loyal", "cheap"]
}`

func Test_Employee(t *testing.T) {
	dat := []byte(EmployeeJSON)
	emp := models.Employee{}

	if err := parsley.Unmarshal(dat, &emp); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if emp.Id != "8becdcce-e762-40cf-b73c-092209f70a30" {
			t.Error("id property value mismatch")
		}
		if emp.Person.Fname != "John" {
			t.Error("person.fname property value mismatch")
		}
		if emp.Person.Lname != "Smith" {
			t.Error("person.lname property value mismatch")
		}
		if emp.Person.DOB.Format(time.RFC3339) != "1998-04-25T00:00:00Z" {
			t.Error("person.dob property value mismatch")
		}
		if len(emp.Devices) != 2 {
			t.Error("person.devices property length mismatch")
		} else {
			if emp.Devices[0].Name != "Work Computer" {
				t.Error("person.devices[0].name property value mismatch")
			}
			if emp.Devices[0].Type != externals.DeviceTypeDesktop {
				t.Error("person.devices[0].type property value mismatch")
			}
			if emp.Devices[1].Name != "Work Mobile" {
				t.Error("person.devices[1].name property value mismatch")
			}
			if emp.Devices[1].Type != externals.DeviceTypeMobile {
				t.Error("person.devices[1].type property value mismatch")
			}
		}
		if emp.IsActive != true {
			t.Error("person.isActive property value mismatch")
		}
		if emp.Rating != 8.65 {
			t.Error("person.rating property value mismatch")
		}
		if emp.LineManager != nil {
			t.Error("person.lineManager property value mismatch")
		}
		if emp.LineManager != nil {
			t.Error("person.lineManager property value mismatch")
		}
		if len(emp.Tags) != 3 {
			t.Error("person.devices property length mismatch")
		} else {
			if emp.Tags[0] != "punctual" {
				t.Error("person.tags[0] element value mismatch")
			}
			if emp.Tags[1] != "loyal" {
				t.Error("person.tags[1] property value mismatch")
			}
			if emp.Tags[2] != "cheap" {
				t.Error("person.tags[2] property value mismatch")
			}
		}
	}
}

const EmployeeListJSON = `[
	{
		"id": "8becdcce-e762-40cf-b73c-092209f70a30",
		"person": {
			"fname": "John",
			"lname": "Smith",
			"dob": "1998-04-25T00:00:00Z"
		},
		"isActive": true,
		"rating": 8.65
	},
	{
		"id": "9d3567d5-f43b-48a3-9d69-7bee5f77776c",
		"person": {
			"fname": "Adam",
			"lname": "Taylor",
			"dob": "1996-02-11T00:00:00Z"
		},
		"isActive": false,
		"rating": 3.05
	}
]`

func Test_EmployeeList(t *testing.T) {
	dat := []byte(EmployeeListJSON)
	emps := models.EmployeeList{}

	if err := parsley.Unmarshal(dat, &emps); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if len(emps) != 2 {
			t.Error("employees length mismatch")
		} else {
			if emps[0].Id != "8becdcce-e762-40cf-b73c-092209f70a30" {
				t.Error("employees[0].id property value mismatch")
			}
			if emps[0].Person.Fname != "John" {
				t.Error("employees[0].person.fname property value mismatch")
			}
			if emps[0].Person.Lname != "Smith" {
				t.Error("employees[0].person.lname property value mismatch")
			}
			if emps[0].Person.DOB.Format(time.RFC3339) != "1998-04-25T00:00:00Z" {
				t.Error("employees[0].person.dob property value mismatch")
			}
			if emps[0].IsActive != true {
				t.Error("employees[0].isActive property value mismatch")
			}
			if emps[0].Rating != 8.65 {
				t.Error("employees[0].rating property value mismatch")
			}

			if emps[1].Id != "9d3567d5-f43b-48a3-9d69-7bee5f77776c" {
				t.Error("employees[1].id property value mismatch")
			}
			if emps[1].Person.Fname != "Adam" {
				t.Error("employees[1].person.fname property value mismatch")
			}
			if emps[1].Person.Lname != "Taylor" {
				t.Error("employees[1].person.lname property value mismatch")
			}
			if emps[1].Person.DOB.Format(time.RFC3339) != "1996-02-11T00:00:00Z" {
				t.Error("employees[1].person.dob property value mismatch")
			}
			if emps[1].IsActive != false {
				t.Error("employees[1].isActive property value mismatch")
			}
			if emps[1].Rating != 3.05 {
				t.Error("employees[1].rating property value mismatch")
			}
		}
	}
}
