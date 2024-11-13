package entity

import (
    "time"
    "gorm.io/gorm"
)

// Entity สำหรับผู้ป่วย (Patient)
type Patient struct {
    gorm.Model
    FirstName      string
    LastName       string
	NationId	   string
    DateOfBirth    time.Time
    Gender         string
    PhoneNumber    string
    Address        string
	HealthInsuranceType    string
	InsurunceId    string
	InsuranceExpiration   time.Time
    MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID"`
    Queues         []Queue         `gorm:"foreignKey:PatientID"`
}

