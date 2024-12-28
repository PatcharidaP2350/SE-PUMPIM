package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicalRecords struct {
	gorm.Model
	SymptomsDetails 	string		`json:"symptoms_details"`
	CheckResults 		string 		`json:"check_results"`
	Diagnose 			string 		`json:"diagnose"`
	OrderMedicine 		string 		`json:"order_medicine"`
	Instructions 		string 		`json:"instructions"`
	Date 				time.Time	`json:"date"`
	Price 				float64		`json:"price"`

	MedicalImages 		[]MedicalImage 	`gorm:"foriegnKey:MedicalRecordsID"`

	SeverityID uint `json:"severity_id"`
	Severity Severity `gorm:"foreignKey:SeverityID"`

	ScheduleID uint	`json:"schedule_id"`
	Schedule Schedule `gorm:"foreignKey:ScheduleID"`

	TakeAHistory *TakeAHistory `gorm:"foreignKey:MedicalRecordsID"`

	Diseases []Disease `gorm:"many2many:medicalrecords_diseases;" json:"diseases"`
	
	EmployeeID uint `json:"employee_id"`
	Employee Employee `gorm:"foreignKey:EmployeeID"`

}
//touch