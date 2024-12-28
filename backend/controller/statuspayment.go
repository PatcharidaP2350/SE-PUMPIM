package controller

import (
	"net/http"
	"SE-B6527075/config"
	"SE-B6527075/entity"
	"github.com/gin-gonic/gin"
)

// ===========================
// Controller สำหรับ StatusPayment
// ===========================

// POST /statuspayments
func CreateStatusPayment(c *gin.Context) {
	var statusPayment entity.StatusPayment

	// Bind JSON เข้าตัวแปร statusPayment
	if err := c.ShouldBindJSON(&statusPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// บันทึก StatusPayment
	if err := db.Create(&statusPayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "StatusPayment created successfully", "data": statusPayment})
}

// GET /statuspayment/:id
func GetStatusPayment(c *gin.Context) {
	ID := c.Param("id")
	var statusPayment entity.StatusPayment

	db := config.DB()
	result := db.First(&statusPayment, ID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	if statusPayment.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, statusPayment)
}

// GET /statuspayments
func ListStatusPayments(c *gin.Context) {
	var statusPayments []entity.StatusPayment

	db := config.DB()
	result := db.Find(&statusPayments)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, statusPayments)
}

// DELETE /statuspayment/:id
func DeleteStatusPayment(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM status_payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "StatusPayment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "StatusPayment deleted successfully"})
}

// PATCH /statuspayment/:id
func UpdateStatusPayment(c *gin.Context) {
	var statusPayment entity.StatusPayment

	StatusPaymentID := c.Param("id")

	db := config.DB()
	result := db.First(&statusPayment, StatusPaymentID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "StatusPayment not found"})
		return
	}

	if err := c.ShouldBindJSON(&statusPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&statusPayment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "StatusPayment updated successfully"})
}