package entity

import (
	"gorm.io/gorm"
)

type Severity struct {
	gorm.Model
	SeverityLevel uint   `json:"severity_level"`
	SeverityName  string `json:"severity_name"`

	MedicalRecords []MedicalRecords `gorm:"foreignKey:SeverityID"`
}

//touch
