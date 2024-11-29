package entity

import (
	"time"

	"gorm.io/gorm"
)

type TakeAHistory struct {
	gorm.Model
	Weight						float32 	`json:"weight"`
	Hight 						float32		`json:"height"`
	PreliminarySymptoms			string		`json:"preliminary_symptoms"`
	SystolicBloodPressure 		uint		`json:"systolic_blood_pressure"`
	DiastolicBloodPressure 		uint		`json:"diastolic_blood_pressure"`
	PulseRate 					uint		`json:"pulse_rate"`
	Smoking 					string		`json:"smoking"`
	LastMenstruationDate 		time.Time	`json:"last_menstruation_date "`

	DrinkAlcohol 				string		`json:"drink_alcohol"`
	Date 						time.Time	`json:"date"`

	MedicalRecordsID *uint           `json:"medical_records_id"`
    MedicalRecords   *MedicalRecords `gorm:"foreignKey:MedicalRecordsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	PatientID uint `json:"patient_id"`
	Patient Patient `gorm:"foriegnKey:PatientID"`

	
	EmployeeID uint `json:"employee_id"`
	Employee Employee `gorm:"foriegnKey:EmployeeID"`
}