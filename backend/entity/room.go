package entity

import "gorm.io/gorm"

type Room struct {

    gorm.Model

    RoomNumber  string `json:"room_number"  valid:"required~room_number is required"`

    RoomTypeID  uint   ` valid:"required~RoomTypeID is required"`
	RoomType  RoomType `gorm:"foreignKey: RoomTypeID"`

    BedCapacity int    `json:"bed_capacity"  valid:"required~bed_capacity is required"`

    DepartmentID uint ` valid:"required~DepartmentID is required"`
    Department Department `gorm:"foreignKey:DepartmentID"`

    EmployeeID uint 
    Employee Employee `gorm:"foreignKey:EmployeeID"`

    PatientRoom []PatientRoom `gorm:"foreignKey:RoomID"`
    
    RoomLayout RoomLayout `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

