package entity

import (
	"time"
	"gorm.io/gorm"
)

type MedicalRecords struct {
    gorm.Model
    SymptomsDetails string        `json:"symptoms_details"`
    CheckResults    string        `json:"check_results"`
    Diagnose        string        `json:"diagnose"`
    OrderMedicine   string        `json:"order_medicine"`
    Instructions    string        `json:"instructions"`
    Date            time.Time     `json:"date"`
    Price           float64       `json:"price"`

    TakeAHistory *TakeAHistory `gorm:"foreignKey:MedicalRecordID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"take_a_history"`

}
