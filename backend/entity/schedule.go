package entity

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	ScheduleName 	string 		`valid:"required~ScheduleName is required" json:"schedule_name"`

	MedicalRecords []MedicalRecords `gorm:"foreignKey:ScheduleID"`
}

//touch