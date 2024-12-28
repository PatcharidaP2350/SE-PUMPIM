package entity

import (
	"time"
	"gorm.io/gorm"
)

type TakeAHistory struct {
	gorm.Model
	Weight						float32 	`json:"weight"`
	Height 						float32		`json:"height"`
	PreliminarySymptoms			string		`json:"preliminary_symptoms"`
	SystolicBloodPressure 		uint		`json:"systolic_blood_pressure"`
	DiastolicBloodPressure 		uint		`json:"diastolic_blood_pressure"`
	PulseRate 					uint		`json:"pulse_rate"`
	Smoking 					bool		`json:"smoking"`
	DrinkAlcohol 				bool		`json:"drink_alcohol"`


	QueueNumber    string    `json:"queue_number"`  
    Date      time.Time    `json:"date"`    
    
    QueueStatus         string    `json:"queue_status"`

	LastMenstruationDate 		time.Time	`json:"last_menstruation_date "`

	PatientVisitID 	uint	`json:"patient_visit_id"`
	PatientVisit PatientVisit `gorm:"foriegnKey:PatientVisitID"`

	MedicalRecordsID *uint           `json:"medical_records_id"`
    MedicalRecords   *MedicalRecords `gorm:"foreignKey:MedicalRecordsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
 
	PatientID uint `json:"patient_id"`
	Patient Patient `gorm:"foriegnKey:PatientID"`
	
	EmployeeID uint `json:"employee_id"`
	Employee Employee `gorm:"foriegnKey:EmployeeID"`

	AppointmentID *uint        `json:"appointment_id"`
    Appointments   *Appointment `gorm:"foreignKey:AppointmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}