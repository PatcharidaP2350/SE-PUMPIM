package entity

import "gorm.io/gorm"

type Hospitalization struct {
	gorm.Model
	HospitalizationName 	string 		`valid:"required~HospitalizationName is required" json:"hospitalization_name"`

	MedicalRecords []MedicalRecords `gorm:"foreignKey:HospitalizationID"`
}

//touch