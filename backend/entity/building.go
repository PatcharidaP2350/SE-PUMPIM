package entity

import "gorm.io/gorm"

type Building struct {
    gorm.Model
    BuildingName  string `json:"building_name" valid:"required~building_name is required"`

    Location      string `json:"location" valid:"required~location is required"`
    
    Floor []Floor `gorm:"foreignKey:BuildingID"`
}
