package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"SE-B6527075/config"
	"SE-B6527075/entity"
)

// GetAllDrugs - ดึงรายการยา
func GetAllDrugs(c *gin.Context) {
    var drugs []entity.Drug

	db := config.DB()
	results := db.Find(&drugs)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, drugs)
	
}


// GetDrugByID - ดึงข้อมูลยาจาก ID
func GetDrugByID(c *gin.Context) {
	db := config.DB()
	id := c.Param("id")
	var drug entity.Drug
	if err := db.First(&drug, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Drug not found"})
		return
	}
	c.JSON(http.StatusOK, drug)
}

// updateDrug 
func UpdateDrug(c *gin.Context) {
	var drug entity.Drug

	drugID := c.Param("id")

	db := config.DB()

	// Find the drug record by ID
	result := db.First(&drug, drugID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Drug ID not found"})
		return
	}

	if err := c.ShouldBindJSON(&drug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&drug)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}

// CreateDrug 
func CreateDrug(c *gin.Context) {
	db := config.DB()

	var input entity.Drug
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
	drug := entity.Drug{
		DrugName:      input.DrugName,
		Category:      input.Category,
		Formulation: input.Formulation,
		Dosage: input.Dosage,
		RegistrationNo: input.RegistrationNo,

		ExpiryDate:    input.ExpiryDate,
		ImportDate:    input.ImportDate,
		StockQuantity: input.StockQuantity,
		ReorderLevel:  input.ReorderLevel,
		PricePerUnit:  input.PricePerUnit,

		Manufacturer: input.Manufacturer,
		CountryOfOrigin: input.CountryOfOrigin,
		BatchNumber: input.BatchNumber,

		SupplierID:    input.SupplierID,

		UsageInstructions: input.UsageInstructions,
		Indications: input.Indications,
		Contraindications: input.Contraindications,
		SideEffects:input.SideEffects,
		Barcode:input.Barcode,
		DrugImage:input.DrugImage,
	}

	// Save drug to database
	if err := db.Create(&drug).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create drug"})
		return
	}

	c.JSON(http.StatusCreated, drug)
}


// DeleteDrug - ลบข้อมูลยา
func DeleteDrug(c *gin.Context) {
	id := c.Param("id")
    db := config.DB()

    result := db.Delete(&entity.Drug{}, id)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
        return
    }
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Drug room with the specified ID not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Drug room deleted successfully"})
}
