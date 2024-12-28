package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicalRecords struct {
	gorm.Model
	SymptomsDetails 	string		`valid:"required~SymptomsDetails is required" json:"symptoms_details"`
	CheckResults 		string 		`valid:"required~CheckResults is required" json:"check_results"`
	Diagnose 			string 		`valid:"required~Diagnose is required" json:"diagnose"`
	OrderMedicine 		string 		`json:"order_medicine"`
	Instructions 		string 		`json:"instructions"`
	Date 				time.Time	`json:"date"`
	Price 				float64		`json:"price"`

	MedicalImages 		[]MedicalImage 	`gorm:"foriegnKey:MedicalRecordsID"`

	SeverityID uint `valid:"required~SeverityID is required" json:"severity_id"`
	Severity Severity `gorm:"foreignKey:SeverityID"`

	ScheduleID uint	`valid:"required~ScheduleID is required" json:"schedule_id"`
	Schedule Schedule `gorm:"foreignKey:ScheduleID"`

	HospitalizationID uint	`valid:"required~HospitalizationID is required" json:"hospitalization_id"`
	Hospitalization Hospitalization `gorm:"foreignKey:HospitalizationID"`

	TakeAHistory *TakeAHistory `gorm:"foreignKey:MedicalRecordsID"`

	Appointment *Appointment `gorm:"foreignKey:MedicalRecordsID"`

	Diseases []Disease `gorm:"many2many:medicalrecords_diseases;" valid:"required~Diseases must have at least one entry" json:"diseases"`
	
	EmployeeID uint `valid:"required~EmployeeID is required" json:"employee_id"`
	Employee Employee `gorm:"foreignKey:EmployeeID"`

}
//touch