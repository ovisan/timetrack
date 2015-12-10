package models

import (
	"time"
)

type Location struct {
	location_code        int    `gorm:"primary_key"`
	location_description string `sql:"size:255"`
}

type Projects struct {
	project_id              int `gorm:"primary_key"`
	location_code           int
	project_mgr_employee_id int
	start_date              time.Time
	end_date                time.Time
	other_details           string `sql:"size:255"`
}

type Employees struct {
	employee_id   int `gorm:"primary_key"`
	start_date    time.Time
	end_date      time.Time
	other_details string `sql:"size:255"`
}

type Activities struct {
	activity_id   int `gorm:"primary_key"`
	activity_code int
	project_id    int
	start_date    time.Time
	end_date      time.Time
	other_details string `sql:"size:255"`
}

type Cost_Centers struct {
	cost_center_id          int    `gorm:"primary_key"`
	cost_center_name        string `sql:"size:255"`
	cost_center_description string `sql:"size:255"`
	other_details           string `sql:"size:255"`
}

type Timesheets struct {
	timesheet_id              int `gorm:"primary_key"`
	activity_id               int
	authorized_by_employee_id int
	cost_center_id            int
	timesheet_for_employee_id int
	start_date                time.Time
	end_date                  time.Time
	submitted_date            time.Time
	other_details             string `sql:"size:255"`
}

func (db *DB) CreateTables() {
	db.CreateTable(&Location{})
	db.CreateTable(&Projects{})
	db.CreateTable(&Employees{})
	db.CreateTable(&Activities{})
	db.CreateTable(&Cost_Centers{})
	db.CreateTable(&Timesheets{})
}
