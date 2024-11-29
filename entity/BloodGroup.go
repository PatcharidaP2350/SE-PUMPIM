package entity

import "gorm.io/gorm"

type BloodGroup struct {
    gorm.Model
    Name string `json:"name"` // เช่น A, B, AB, O
	RHFactor string `json:"rh_factor"` // เช่น Rh+ หรือ Rh-
}
