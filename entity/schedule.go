package entity

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	ScheduleName 	string 		`json:"schedule_name"`

	MedicalRecords []MedicalRecords `gorm:"foreignKey:ScheduleID"`
}

//touch