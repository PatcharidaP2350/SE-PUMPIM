package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"SE-B6527075/config"
	"SE-B6527075/entity"
)

// GetAllDrugs - ดึงรายการยา
func GetAllInventory(c *gin.Context) {
	var inventory []entity.Inventory

	db := config.DB()
	results := db.Find(&inventory)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// GetDrugByID - ดึงข้อมูลยาจาก ID
func GetInventoryByID(c *gin.Context) {
	db := config.DB()

	id := c.Param("id")
	var inventory entity.Inventory
	if err := db.First(&inventory, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "inventory not found"})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// updateDrug
func UpdateInventory(c *gin.Context) {
	var inventory entity.Inventory

	inventoryID := c.Param("id")

	db := config.DB()
	result := db.First(&inventory, inventoryID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&inventory)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}

// CreateDrug
func CreateInventory(c *gin.Context) {
	db := config.DB()

	var input entity.Inventory
	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if supplier exists
	var drug entity.Drug
	if err := db.First(&drug, input.DrugID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":        "drug not found",
			"drug_id":  input.DrugID,
		})
		return
	}
	var medicalEquipment entity.MedicalEquipment
	if err := db.First(&medicalEquipment, input.MedicalEquipmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":        "medicalEquipment not found",
			"medicalEquipment_id":  input.MedicalEquipmentID,
		})
		return
	}

	// Create new drug
	inventory := entity.Inventory{
		Location:      input.Location,
		Quantity: input.Quantity,
		MedicalEquipmentID:    input.MedicalEquipmentID,
		DrugID: input.DrugID,

	}

	// Save drug to database
	if err := db.Create(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory"})
		return
	}

	c.JSON(http.StatusCreated, inventory)
}

// DeleteDrug - ลบข้อมูลยา
func DeleteInventory(c *gin.Context) {
	db := config.DB()

	id := c.Param("id")
	if err := db.Delete(&entity.Drug{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete drug"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Drug deleted successfully"})
}
