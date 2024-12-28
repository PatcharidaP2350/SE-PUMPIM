package entity

import (
	"gorm.io/gorm"
)

type Severity struct {
	gorm.Model
	SeverityLevel uint   `valid:"required~SeverityLevel is required" json:"severity_level"`
	SeverityName  string `valid:"required~SeverityName is required" json:"severity_name"`

	MedicalRecords []MedicalRecords `gorm:"foreignKey:SeverityID"`
}

//touch
