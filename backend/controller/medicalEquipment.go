package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"SE-B6527075/config"
	"SE-B6527075/entity"
)


func GetAllMedicalEquipment(c *gin.Context){
	var medicalEquipment []entity.MedicalEquipment

	db := config.DB()

	results := db.Find(&medicalEquipment)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, medicalEquipment)
}

func GetMedicalEquipmentByID(c *gin.Context){
	var medicalequipment entity.MedicalEquipment

	db := config.DB()

	id := c.Param("id")

	if err := db.First(&medicalequipment, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "medicalequipment not found"})
		return
	}
	c.JSON(http.StatusOK, medicalequipment)
}

func UpdateMedicalEquipment(c *gin.Context) {
	db := config.DB()

	var medicalEquipment entity.MedicalEquipment
	// แปลง JSON ที่ส่งมาเป็น struct
	if err := c.ShouldBindJSON(&medicalEquipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// ตรวจสอบว่ามี ID หรือไม่ (จำเป็นสำหรับการอัปเดต)
	if medicalEquipment.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required for updating"})
		return
	}

	// ใช้ db.Model().Updates() เพื่ออัปเดตข้อมูลที่มีอยู่
	if err := db.Model(&entity.MedicalEquipment{}).Where("id = ?", medicalEquipment.ID).Updates(medicalEquipment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update medicalEquipment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "MedicalEquipment updated successfully", "data": medicalEquipment})
}


func DeleteMedicalEquipment(c *gin.Context) {
    id := c.Param("id")
    db := config.DB()

    result := db.Delete(&entity.MedicalEquipment{}, id)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
        return
    }
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Equipment room with the specified ID not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Equipment room deleted successfully"})
}

  

func CreateMedicalEquiment(c *gin.Context) {
	db := config.DB()

	var input entity.MedicalEquipment
	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if supplier exists
	var supplier entity.Supplier
	if err := db.First(&supplier, input.SupplierID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":        "Supplier not found",
			"supplier_id":  input.SupplierID,
		})
		return
	}

	// Create new drug
	medicalEquipment := entity.MedicalEquipment{
		Name:      input.Name,
		Category:      input.Category,
		EquipmentModel:    input.EquipmentModel,
		SerialNumber: input.SerialNumber,
		Manufacturer: input.Manufacturer,
		CountryOfOrigin:input.CountryOfOrigin,

		
		StockQuantity: input.StockQuantity,
		ReorderLevel: input.ReorderLevel,
		PricePerUnit: input.PricePerUnit,
		ImportDate:    input.ImportDate,
		ExpiryDate: input.ExpiryDate,

		LastMaintenance: input.LastMaintenance,
		MaintenanceSchedule: input.MaintenanceSchedule,
		MaintenanceHistory: input.MaintenanceHistory,

		SupplierID:    input.SupplierID,

		UsageInstructions:input.UsageInstructions,
		Warnings:input.Warnings,
		Purpose:input.Purpose,
		Image:input.Image,
		Standards:input.Standards,
	}

	// Save drug to database
	if err := db.Create(&medicalEquipment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create medicalEquipment"})
		return
	}

	c.JSON(http.StatusCreated, medicalEquipment)
}