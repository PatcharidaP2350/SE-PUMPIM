package entity

import (
	"gorm.io/gorm"
)

type Specialist struct {
	gorm.Model
	SpecialistName string     `valid:"required~SpecialistName is required" json:"specialist_name"`
	Employees      []Employee `gorm:"foreignKey:SpecialistID"` // เชื่อมกับ Employee โดยใช้ SpecialistID เป็น foreign key
}
