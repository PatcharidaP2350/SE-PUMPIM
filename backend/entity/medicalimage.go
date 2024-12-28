package entity

import "gorm.io/gorm"

type MedicalImage struct {
	gorm.Model
	Image 	string 		`json:"image" gorm:"type:longtext"`

	MedicalRecordsID 	uint `json:"medical_records_id"`
	MedicalRecords 		MedicalRecords `gorm:"foreignKey:MedicalRecordsID"`
}

//touch