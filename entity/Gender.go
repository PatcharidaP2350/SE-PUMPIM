package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	GenderName string     `json:"gender_name"`
	Patients  []Patient `gorm:"foreignKey:GenderID"`
	Employees  []Employee `gorm:"foreignKey:GenderID"` // เชื่อมกับ Employee โดยใช้ GenderID เป็น foreign key
}