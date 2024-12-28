package entity

import (
	"gorm.io/gorm"
)

type PatientVisit struct {
	gorm.Model
	PatientVisitType string     `json:"patient_visit_type"`
	
	TakeAHistory  []TakeAHistory `gorm:"foreignKey:PatientVisitID"`
}