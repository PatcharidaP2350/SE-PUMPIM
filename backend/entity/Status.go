package entity

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	StatusName string     `valid:"required~StatusName is required" json:"status_name"`
	Employees  []Employee `gorm:"foreignKey:StatusID"` // เชื่อมกับ Employee โดยใช้ StatusID เป็น foreign key
}
