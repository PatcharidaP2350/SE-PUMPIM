package entity

import (
    "gorm.io/gorm"
    "time"
)

type Appointment struct {
    gorm.Model
    AppointmentDate time.Time      `json:"appoint_date"`
    AppointmentTime       time.Time       `json:"appointment_time"`
    Reason        string       `json:"reason"`
    Status      time.Time    `json:"status"`

    TakeAHistory *TakeAHistory `gorm:"foreignKey:AppointmentID"`
    
	MedicalRecordsID *uint           `json:"medical_records_id"`
    MedicalRecords   *MedicalRecords `gorm:"foreignKey:MedicalRecordsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    
	EmployeeID *uint `json:"employee_id"`
	Employee *Employee `gorm:"foriegnKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
 
	// Appointment *Appointment `gorm:"foreignKey:AppointmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"employee"`  //ใส่ที่ Employee
} 
