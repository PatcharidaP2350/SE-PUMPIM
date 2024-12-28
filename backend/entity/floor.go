package entity

import "gorm.io/gorm"

type Floor struct {

	gorm.Model

    FloorNumber     string    `json:"floor_number" valid:"required~floor_number is required"`

	BuildingID uint `valid:"required~BuildingID is required"`
	Building Building `gorm:"foreignKey:BuildingID" `

	RoomLayout []RoomLayout `gorm:"foreignKey:FloorID"`
}
