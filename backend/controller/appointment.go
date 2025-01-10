package controller

import (
	"fmt"
	"net/http"
	"time"

	"SE-B6527075/config"
	"SE-B6527075/entity"

	"github.com/gin-gonic/gin"
)

func CreateAppointment(c *gin.Context){
	fmt.Println("Creating or Updating TakeAHistory")
	db := config.DB()

	// กำหนดข้อมูลที่รับจาก Request
	var input struct {
		AppointmentDate             time.Time    `json:"appoint_date"`
		AppointmentTime             time.Time    `json:"appointment_time"`
		Reason                      string       `json:"reason"`
		Status                      string    `json:"status"`
		Note 						string 	  `json:"note"`
		EmployeeID                  uint         `json:"employee_id" binding:"required"`
		MedicalRecordsID            *uint        `json:"medical_records_id"`
	}

	// ดึงข้อมูล JSON จากคำขอ (Request) และตรวจสอบความถูกต้อง
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	// สร้าง record ใหม่จากข้อมูลที่ได้รับ
	app := entity.Appointment{
		AppointmentDate:             input.AppointmentDate,
		AppointmentTime:             input.AppointmentTime,
		Reason:                      input.Reason,
		Status:                      input.Status,
		Note:                        input.Note,
		EmployeeID:                  input.EmployeeID,
		MedicalRecordsID:            input.MedicalRecordsID,
	}

	// เริ่มต้นการเชื่อมต่อฐานข้อมูล
	if err := db.Create(&app).Error; err != nil {
		fmt.Println("Error creating appointment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	// ส่งการตอบกลับที่ประสบความสำเร็จ
	c.JSON(http.StatusCreated, gin.H{
		"message": "Appointment created and TakeAHistory handled successfully",
		"data":    app,
	})
}

func GetAppointmentByID(c *gin.Context) {
	// รับค่า AppointmentID จากพารามิเตอร์ URL
    id := c.Param("id")

    // ตัวแปรสำหรับเก็บข้อมูล TakeAHistory
    var appointment entity.Appointment

    // ดึงข้อมูลจากฐานข้อมูลโดยใช้ ID
    db := config.DB()
    if err := db.Preload("MedicalRecords").Preload("Employee").First(&appointment, "id = ?", id).Error; err != nil {
        // ถ้าไม่พบข้อมูล ให้ส่งข้อความ error กลับ
        c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
        return
    }

    // ส่งข้อมูลกลับในรูปแบบ JSON
    c.JSON(http.StatusOK, gin.H{"data": appointment})
}