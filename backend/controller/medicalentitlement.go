package controller

import (
	"fmt"
	"net/http"

	"SE-B6527075/config"
	"SE-B6527075/entity"
	"github.com/gin-gonic/gin"
)

// CreateMedicalEntitlement handles creating a new Medical Entitlement
func CreateMedicalEntitlement(c *gin.Context) {
	var medicalEntitlement entity.MedicalEntitlement

	// Bind JSON
	if err := c.ShouldBindJSON(&medicalEntitlement); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// Check for duplicate MedicalEntitlementName
	var existingMedicalEntitlement entity.MedicalEntitlement
	if err := db.Where("medical_entitlement_name = ?", medicalEntitlement.MedicalEntitlementName).First(&existingMedicalEntitlement).Error; err == nil {
		fmt.Println("Medical Entitlement already exists:", medicalEntitlement.MedicalEntitlementName)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medical Entitlement already exists"})
		return
	}

	// Save to database
	if err := db.Create(&medicalEntitlement).Error; err != nil {
		fmt.Println("Error creating Medical Entitlement:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Medical Entitlement created successfully:", medicalEntitlement)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created success",
		"data":    medicalEntitlement,
	})
}

// GetMedicalEntitlement handles fetching a single Medical Entitlement by ID
func GetMedicalEntitlement(c *gin.Context) {
	ID := c.Param("id")
	var medicalEntitlement entity.MedicalEntitlement

	db := config.DB()
	if err := db.First(&medicalEntitlement, ID).Error; err != nil {
		fmt.Println("Medical Entitlement not found, ID:", ID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Medical Entitlement not found"})
		return
	}

	c.JSON(http.StatusOK, medicalEntitlement)
}

// ListMedicalEntitlements handles fetching all Medical Entitlements
func ListMedicalEntitlements(c *gin.Context) {
	var medicalEntitlements []entity.MedicalEntitlement

	db := config.DB()
	if err := db.Find(&medicalEntitlements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicalEntitlements)
}

// UpdateMedicalEntitlement handles updating an existing Medical Entitlement
func UpdateMedicalEntitlement(c *gin.Context) {
	ID := c.Param("id")
	var medicalEntitlement entity.MedicalEntitlement

	db := config.DB()
	if err := db.First(&medicalEntitlement, ID).Error; err != nil {
		fmt.Println("Medical Entitlement not found, ID:", ID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Medical Entitlement not found"})
		return
	}

	var updatedData entity.MedicalEntitlement
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for duplicate MedicalEntitlementName (exclude the current one)
	var otherMedicalEntitlement entity.MedicalEntitlement
	if err := db.Where("medical_entitlement_name = ? AND id != ?", updatedData.MedicalEntitlementName, ID).First(&otherMedicalEntitlement).Error; err == nil {
		fmt.Println("Medical Entitlement name already exists:", updatedData.MedicalEntitlementName)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medical Entitlement name already exists"})
		return
	}

	// Update fields
	medicalEntitlement.MedicalEntitlementName = updatedData.MedicalEntitlementName
	medicalEntitlement.MedicalEntitlementType = updatedData.MedicalEntitlementType
	medicalEntitlement.MedicalEntitlementUsageLimit = updatedData.MedicalEntitlementUsageLimit
	medicalEntitlement.MedicalEntitlementProviderName = updatedData.MedicalEntitlementProviderName
	medicalEntitlement.MedicalEntitlementProviderContact = updatedData.MedicalEntitlementProviderContact

	if err := db.Save(&medicalEntitlement).Error; err != nil {
		fmt.Println("Error updating Medical Entitlement:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Medical Entitlement updated successfully:", medicalEntitlement)
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated success",
		"data":    medicalEntitlement,
	})
}

// DeleteMedicalEntitlement handles deleting a Medical Entitlement by ID
func DeleteMedicalEntitlement(c *gin.Context) {
	ID := c.Param("id")
	db := config.DB()

	if tx := db.Delete(&entity.MedicalEntitlement{}, ID); tx.RowsAffected == 0 {
		fmt.Println("Medical Entitlement not found, ID:", ID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medical Entitlement not found"})
		return
	}

	fmt.Println("Medical Entitlement deleted successfully, ID:", ID)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted success"})
}
