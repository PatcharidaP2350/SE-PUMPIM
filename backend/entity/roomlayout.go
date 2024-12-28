package entity
import (
	"gorm.io/gorm"
)

type RoomLayout struct{
	gorm.Model

	RoomID uint `valid:"required~RoomID is required"`
	Room   *Room `gorm:"foreignKey:RoomID"`

	FloorID uint `valid:"required~FloorID is required"`
	Floor Floor `json:"foreignKey:FloorID"`

	PositionX int `json:"position_x" valid:"required~position_x is required"`

	PositionY int `json:"position_y" valid:"required~position_y is required"`
}