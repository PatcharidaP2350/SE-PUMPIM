package entity

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name        string        // ชื่อบริษัท/ผู้จัดจำหน่าย
    Contact     string        // ข้อมูลการติดต่อ
    Address     string
	PhoneNumber string
	Email       string
}