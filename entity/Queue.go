package entity

import (
    "gorm.io/gorm"
    "time"
)
type Queue struct {
    gorm.Model
    QueueNumber    string    `json:"queue_number"`  
    QueueDate      string    `json:"queue_date"`    
    QueueTime      time.Time    `json:"queue_time"`    
    Status         string    `json:"status"`        
    PatientID      uint      `json:"patient_id"` 
    Patient        Patient   `gorm:"foreignKey:PatientID" ` 
}
