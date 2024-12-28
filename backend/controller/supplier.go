package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"SE-B6527075/config"
	"SE-B6527075/entity"
)

// GetAllDrugs - ดึงรายการยา
func GetAllSuppliers(c *gin.Context) {
	var suppliers []entity.Supplier
	db := config.DB()

	result := db.Find(&suppliers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}


func GetSupplierByID(c *gin.Context) {
	var supplier entity.Supplier
	db := config.DB()

	id := c.Param("id")
	if err := db.First(&supplier, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	c.JSON(http.StatusOK, supplier)
}



// CreateDrug 
func CreateSupplier(c *gin.Context) {
	db := config.DB()

	var input entity.Supplier
	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	

	// Create new drug
	supplier := entity.Supplier{
		Name:      input.Name,
		PhoneNumber:    input.PhoneNumber,
		Email: input.Email,
		Address: input.Address,
	}

	// Save drug to database
	if err := db.Create(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create supplier"})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}
// DeleteDrug - ลบข้อมูลยา
func UpdateSupplier(c *gin.Context) {
	var supplier entity.Supplier
	db := config.DB()

	// ดึงข้อมูล Supplier ที่ต้องการแก้ไขจาก ID
	id := c.Param("id")
	if err := db.First(&supplier, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	// อัปเดตข้อมูลจาก Body ของ Request
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึกการเปลี่ยนแปลงลงฐานข้อมูล
	if err := db.Save(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

func DeleteSupplier(c *gin.Context) {
	var supplier entity.Supplier
	db := config.DB()

	// ดึงข้อมูล Supplier ที่ต้องการลบจาก ID
	id := c.Param("id")
	if err := db.First(&supplier, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	// ลบข้อมูลออกจากฐานข้อมูล
	if err := db.Delete(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully"})
}

