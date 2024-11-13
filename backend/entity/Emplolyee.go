package entity

import (
	"time"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName          string
	LastName           string
	Age                uint
	DateOfBirth        time.Time
	Email              string
	Phone              string
	Address            string
	Username           string  
	ProfessionalLicense string 
	CongenitalDisease  string  
	Graduate           string  

	// Foreign Keys and Relationships
	// GenderID     uint       
	// Gender       Gender     `gorm:"foreignKey:GenderID"`

	// PositionID   uint       
	// Position     Position   `gorm:"foreignKey:PositionID"`

	// DepartmentID uint       
	// Department   Department `gorm:"foreignKey:DepartmentID"`

	// StatusID     uint       
	// Status       Status     `gorm:"foreignKey:StatusID"`

	// SpecialistID uint       
	// Specialist   Specialist `gorm:"foreignKey:SpecialistID"`

	// Profile      string     `gorm:"type:longtext"`
	// Password     string

	// New Relationship with MedicalRecord
	MedicalRecords []MedicalRecord `gorm:"foreignKey:DoctorID"`  // One-to-many relationship with MedicalRecord
}
