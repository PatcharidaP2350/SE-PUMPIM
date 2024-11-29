package entity

import (
	"time"
	"gorm.io/gorm"
)

type TakeAHistory struct {
    gorm.Model
    Weight               float32       `json:"weight"`
    Height               float32       `json:"height"`
    PreliminarySymptoms  string        `json:"preliminary_symptoms"`
    SystolicBloodPressure uint         `json:"systolic_blood_pressure"`
    DiastolicBloodPressure uint        `json:"diastolic_blood_pressure"`
    PulseRate            uint          `json:"pulse_rate"`
    Smoking              string        `json:"smoking"`
    DrinkAlcohol         string        `json:"drink_alcohol"`
    Date                 time.Time     `json:"date"`

    // Foreign Key เชื่อมกับ MedicalRecords
    MedicalRecordID *uint           `json:"medical_record_id"`
    MedicalRecord   *MedicalRecords `gorm:"foreignKey:MedicalRecordID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

}
