package controller

import (
	"net/http"
	"SE-B6527075/config"
	"SE-B6527075/entity"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
)

// POST /payments
func CreatePayment(c *gin.Context) {
	var payment entity.Payment
	var input struct {
		entity.Payment
		PaymentServiceID string `json:"payment_service_id"`
		PaymentMethod    uint `json:"payment_method"`
		PaymentNotes     string `json:"payment_notes"`
	}

	// Bind JSON เข้าตัวแปร input
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// ตรวจสอบว่า PaymentServiceID ซ้ำกันหรือไม่
	var existingPayment entity.Payment
	if err := db.Where("payment_service_id = ?", input.PaymentServiceID).First(&existingPayment).Error; err == nil {
		// ถ้าเจอ Payment ที่มี PaymentServiceID ซ้ำ
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentServiceID นี้ถูกใช้งานแล้ว"})
		return
	}

	// ตรวจสอบว่า MedicalRecordsID ซ้ำกันหรือไม่
	var existingMedicalRecords entity.Payment
	if err := db.Where("medical_records_id = ?", input.MedicalRecordsID).First(&existingMedicalRecords).Error; err == nil {
		// ถ้าเจอ Payment ที่มี MedicalRecordsID ซ้ำ
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalRecordsID นี้ถูกใช้งานแล้ว"})
		return
	}

	// ตรวจสอบว่า Patient, Employee, MedicalRecords, StatusPayment ถูกต้องหรือไม่
	var patient entity.Patient
	if err := db.First(&patient, input.PatientID).Error; err != nil {
		fmt.Println("Patient not found, ID:", input.PatientID)
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบผู้ป่วย"})
		return
	}

	var employee entity.Employee
	if err := db.First(&employee, input.EmployeeID).Error; err != nil {
		fmt.Println("Employee not found, ID:", input.EmployeeID)
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบพนักงาน"})
		return
	}

	var medicalRecords entity.MedicalRecords
	if err := db.First(&medicalRecords, input.MedicalRecordsID).Error; err != nil {
		fmt.Println("MedicalRecords not found, ID:", input.MedicalRecordsID)
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลทางการแพทย์"})
		return
	}

	var statusPayment entity.StatusPayment
	if err := db.First(&statusPayment, input.StatusPaymentID).Error; err != nil {
		fmt.Println("StatusPayment not found, ID:", input.StatusPaymentID)
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบสถานะการชำระเงิน"})
		return
	}

	// กำหนดค่าฟิลด์ใน Payment
	payment.PatientID = input.PatientID
	payment.EmployeeID = input.EmployeeID
	payment.MedicalRecordsID = input.MedicalRecordsID
	payment.StatusPaymentID = input.StatusPaymentID
	payment.PaymentServiceID = input.PaymentServiceID
	payment.PaymentMethodID = input.PaymentMethod
	payment.PaymentNotes = input.PaymentNotes
	payment.PaymentDate = input.PaymentDate
	payment.PaymentAmount = input.PaymentAmount

	// บันทึก Payment ลงฐานข้อมูล
	if err := db.Create(&payment).Error; err != nil {
		fmt.Println("Error creating payment:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// โหลดข้อมูลที่เกี่ยวข้อง (Preload)
	if err := db.Preload("Patient").
		Preload("Employee").
		Preload("MedicalRecords").
		Preload("StatusPayment").
		First(&payment, payment.ID).Error; err != nil {
		fmt.Println("Error loading related data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load related data"})
		return
	}

	// ส่งข้อมูลกลับมาพร้อม status 201
	c.JSON(http.StatusCreated, gin.H{
		"message": "Payment created successfully", // ส่งข้อความบ่งบอกว่าการสร้างสำเร็จ
		"data":    payment,                      // ส่งข้อมูล Payment ที่ถูกสร้าง
	})
}

func CreatePaymentByIDMedicalID(c *gin.Context) {
	var input struct {
		TakeHistoryID uint    `json:"takehistory_id"`
		Price         float64 `json:"price"`
	}

	// Bind JSON เข้าไปใน input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// ค้นหาค่า patient_id จาก takehistory_id
	var takeHistory entity.TakeAHistory
	if err := db.First(&takeHistory, input.TakeHistoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบ TakeHistory ที่ระบุ"})
		return
	}

	// หา MedicalRecords ที่ถูกสร้างล่าสุด
	var medicalRecord entity.MedicalRecords
	if err := db.Order("created_at desc").First(&medicalRecord).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบ MedicalRecords ที่สร้างล่าสุด"})
		return
	}

	// ใช้ employee_id จาก medicalRecord
	var employee entity.Employee
	if err := db.First(&employee, medicalRecord.EmployeeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบ Employee ที่เกี่ยวข้อง"})
		return
	}

	// สร้าง Payment ใหม่
	payment := entity.Payment{
		PatientID:        takeHistory.PatientID,
		EmployeeID:       employee.ID,
		MedicalRecordsID: medicalRecord.ID,
		StatusPaymentID:  1, // ค่าสถานะการชำระเงิน (ใช้ค่านี้ตามที่กำหนดหรือค่อยอัพเดต)
		PaymentServiceID: "PS123456", // ตัวอย่าง PaymentServiceID
		PaymentMethodID:    1, // ตัวอย่าง PaymentMethod
		PaymentNotes:     "Payment for medical treatment", // ตัวอย่าง Payment Notes
		PaymentDate:      time.Now(),
		PaymentAmount:    input.Price,
	}

	// บันทึก Payment ลงฐานข้อมูล
	if err := db.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้าง Payment ได้"})
		return
	}

	// ส่งข้อมูลกลับมา
	c.JSON(http.StatusCreated, gin.H{
		"message": "Payment created successfully",
		"data":    payment,
	})
}





// GET /payment/:id
func GetPayment(c *gin.Context) {
	ID := c.Param("id")
	var payment entity.Payment

	db := config.DB()

	// ใช้ Preload ดึงข้อมูลจากตาราง MedicalEntitlements ผ่านตารางกลาง
	result := db.Preload("MedicalEntitlements").First(&payment, ID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// ตรวจสอบว่า Payment ถูกพบหรือไม่
	if payment.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, payment)
}


// GET /payments
func ListPayments(c *gin.Context) {
	var payments []entity.Payment

	db := config.DB()

	// ใช้ Preload เพื่อดึงข้อมูลจากตาราง MedicalEntitlements
	result := db.Preload("MedicalEntitlements").Find(&payments)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}


func ListPaymentsNoPay(c *gin.Context) {
	var payments []entity.Payment

	db := config.DB()

	// ใช้ Preload เพื่อดึงข้อมูลจากตาราง MedicalEntitlements พร้อมตั้งเงื่อนไข StatusPaymentID = 1
	result := db.Preload("MedicalEntitlements").Where("status_payment_id = ?", 2).Find(&payments)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}


// DELETE /payment/:id
func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}

// PATCH /payment/:id
func UpdatePayment(c *gin.Context) {
	var payment entity.Payment

	PaymentID := c.Param("id")

	db := config.DB()
	result := db.First(&payment, PaymentID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&payment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully"})
}
