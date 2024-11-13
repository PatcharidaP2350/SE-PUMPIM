package entity

import (
    "time"
    "gorm.io/gorm"
)

// Entity สำหรับคิวผู้ป่วย (Queue)
type Queue struct {
    gorm.Model
    QueueNumber int       // หมายเลขคิว
    Status      string    // สถานะ เช่น "Waiting", "In Progress", "Completed"
    QueueDate   time.Time // วันที่และเวลาของคิว
    PatientID   uint      // FK ไปยัง Patient (ผู้ป่วยที่อยู่ในคิว)
}