package entity

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Location          string
    DrugID            uint   
    MedicalEquipmentID uint   
    Quantity          int  
}