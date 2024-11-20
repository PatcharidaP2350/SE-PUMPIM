package entity

import (
    "gorm.io/gorm"
    "time"
)

type Patient struct {
    gorm.Model
    IdentificationNumber string      `json:"identification_number"`
    FirstName       string       `json:"first_name"`
    LastName        string       `json:"last_name"`
    DateOfBirth      time.Time    `json:"date_of_birth"`
    Address         string       `json:"address"`          
    PhoneNumber     string       `json:"phone_number"`     
    GenderID        uint         `json:"gender_id"`
    BloodGroupID    uint         `json:"blood_group_id"`
    // PatientDiseases []PatientDisease `gorm:"foreignKey:PatientID" `
    Diseases []Disease `gorm:"many2many:patient_diseases;" json:"diseases"`

    Gender     Gender     `gorm:"foreignKey:GenderID"`
    BloodGroup BloodGroup `gorm:"foreignKey:BloodGroupID"`
}