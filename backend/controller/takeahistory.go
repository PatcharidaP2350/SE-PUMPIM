package controller

import (
	"net/http"
	"time"
	"fmt"

	"SE-B6527075/entity"
	"SE-B6527075/config"
	"github.com/gin-gonic/gin"
)




func CreateTakeAHistory(c *gin.Context) {
	fmt.Println("Creating or Updating Medical Record")
	db := config.DB()

	// กำหนดข้อมูลที่รับจาก Request
	var input struct {
		Weight                   float32    `json:"weight" binding:"required"`
		Hight                    float32    `json:"hight" binding:"required"`
		PreliminarySymptoms      string     `json:"preliminary_symptoms" binding:"required"`
		SystolicBloodPressure    uint       `json:"systolic_blood_pressure" binding:"required"`
		DiastolicBloodPressure   uint       `json:"diastolic_blood_pressure" binding:"required"`
		PulseRate                uint       `json:"pulse_rate" binding:"required"`
		Smoking                  string     `json:"smoking" binding:"required"`
		LastMenstruationDate     time.Time  `json:"last_menstruation_date" binding:"required"`
		DrinkAlcohol             string     `json:"drink_alcohol" binding:"required"`
		PatientID                uint       `json:"patient_id" binding:"required"`
		EmployeeID               uint       `json:"employee_id" binding:"required"`
	}

	// ดึงข้อมูล JSON จากคำขอ (Request) และตรวจสอบความถูกต้อง
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	// สร้าง record ใหม่จากข้อมูลที่ได้รับ
	take := entity.TakeAHistory{
		Weight:                  input.Weight,
		Hight:                   input.Hight,
		PreliminarySymptoms:     input.PreliminarySymptoms,
		SystolicBloodPressure:   input.SystolicBloodPressure,
		DiastolicBloodPressure:  input.DiastolicBloodPressure,
		PulseRate:               input.PulseRate,
		Smoking:                 input.Smoking,
		LastMenstruationDate:    input.LastMenstruationDate,
		DrinkAlcohol:            input.DrinkAlcohol,
		Date:                    time.Now(),
		MedicalRecordsID:       nil, // ปรับเปลี่ยนตามความเหมาะสม
		PatientID:               input.PatientID,
		EmployeeID:              input.EmployeeID,
	}

	// เริ่มต้นการเชื่อมต่อฐานข้อมูล
	if err := db.Create(&take).Error; err != nil {
		fmt.Println("Error saving medical record:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save medical record", "details": err.Error()})
		return
	}

	// ส่งการตอบกลับที่ประสบความสำเร็จ
	c.JSON(http.StatusCreated, gin.H{
		"message": "Medical record created and TakeAHistory handled successfully",
		"data":    take,
	})
}


func ListTakeAHistory(c *gin.Context) {
    db := config.DB() // เรียกใช้การเชื่อมต่อฐานข้อมูลจาก config

    var takeAHistories []entity.TakeAHistory

    // ดึงข้อมูลทั้งหมดจากฐานข้อมูล พร้อมกับ preload ความสัมพันธ์ที่เกี่ยวข้อง (เช่น Patient, Employee หรือ MedicalRecords ถ้ามี)
    if err := db.Preload("Patient").Preload("Employee").Preload("MedicalRecords").Find(&takeAHistories).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve TakeAHistory", "details": err.Error()})
        return
    }

    // ส่งข้อมูลกลับในรูปแบบ JSON
    c.JSON(http.StatusOK, gin.H{
        "message": "TakeAHistory retrieved successfully",
        "data":    takeAHistories,
    })
}


func GetTakeAHistory(c *gin.Context) {
    // รับค่า TakeAHistoryID จากพารามิเตอร์ URL
    id := c.Param("id")

    // ตัวแปรสำหรับเก็บข้อมูล TakeAHistory
    var takeAHistory entity.TakeAHistory

    // ดึงข้อมูลจากฐานข้อมูลโดยใช้ ID
    db := config.DB()
    if err := db.Preload("Patient").Preload("Employee").First(&takeAHistory, "id = ?", id).Error; err != nil {
        // ถ้าไม่พบข้อมูล ให้ส่งข้อความ error กลับ
        c.JSON(http.StatusNotFound, gin.H{"error": "TakeAHistory not found"})
        return
    }

    // ส่งข้อมูลกลับในรูปแบบ JSON
    c.JSON(http.StatusOK, gin.H{"data": takeAHistory})
}


func UpdateTakeAHistory(c *gin.Context) {
	fmt.Println("Updating TakeAHistory Record")
	db := config.DB()

	// รับ TakeAHistoryID จาก URL Parameter
	id := c.Param("id")

	// ตรวจสอบว่ามี TakeAHistory ที่ต้องการอัปเดตหรือไม่
	var take entity.TakeAHistory
	if err := db.First(&take, id).Error; err != nil {
		fmt.Println("Record not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "TakeAHistory record not found"})
		return
	}

	// กำหนดโครงสร้างข้อมูลที่รับจาก Request
	var input struct {
		Weight                   *float32   `json:"weight"`
		Hight                    *float32   `json:"hight"`
		PreliminarySymptoms      *string    `json:"preliminary_symptoms"`
		SystolicBloodPressure    *uint      `json:"systolic_blood_pressure"`
		DiastolicBloodPressure   *uint      `json:"diastolic_blood_pressure"`
		PulseRate                *uint      `json:"pulse_rate"`
		Smoking                  *string    `json:"smoking"`
		LastMenstruationDate     *time.Time `json:"last_menstruation_date"`
		DrinkAlcohol             *string    `json:"drink_alcohol"`
	}

	// ดึงข้อมูล JSON จากคำขอ (Request) และตรวจสอบความถูกต้อง
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	// อัปเดตฟิลด์ที่ส่งมาใน Request
	if input.Weight != nil {
		take.Weight = *input.Weight
	}
	if input.Hight != nil {
		take.Hight = *input.Hight
	}
	if input.PreliminarySymptoms != nil {
		take.PreliminarySymptoms = *input.PreliminarySymptoms
	}
	if input.SystolicBloodPressure != nil {
		take.SystolicBloodPressure = *input.SystolicBloodPressure
	}
	if input.DiastolicBloodPressure != nil {
		take.DiastolicBloodPressure = *input.DiastolicBloodPressure
	}
	if input.PulseRate != nil {
		take.PulseRate = *input.PulseRate
	}
	if input.Smoking != nil {
		take.Smoking = *input.Smoking
	}
	if input.LastMenstruationDate != nil {
		take.LastMenstruationDate = *input.LastMenstruationDate
	}
	if input.DrinkAlcohol != nil {
		take.DrinkAlcohol = *input.DrinkAlcohol
	}

	// บันทึกการเปลี่ยนแปลงลงฐานข้อมูล
	if err := db.Save(&take).Error; err != nil {
		fmt.Println("Error updating record:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update TakeAHistory record", "details": err.Error()})
		return
	}

	// ส่งการตอบกลับเมื่ออัปเดตสำเร็จ
	c.JSON(http.StatusOK, gin.H{
		"message": "TakeAHistory record updated successfully",
		"data":    take,
	})
}


func DeleteTakeAHistory(c *gin.Context) {
    // รับค่า TakeAHistoryID จากพารามิเตอร์ URL
    id := c.Param("id")

	db := config.DB()

    // ตัวแปรสำหรับเก็บข้อมูล TakeAHistory ที่จะลบ
    var takeAHistory entity.TakeAHistory

    // ค้นหาข้อมูล TakeAHistory โดยใช้ ID
    if err := db.First(&takeAHistory, id).Error; err != nil {
        // ถ้าไม่พบข้อมูล ให้ส่งข้อความ error กลับ
        c.JSON(http.StatusNotFound, gin.H{"error": "TakeAHistory not found"})
        return
    }

    // ลบข้อมูล TakeAHistory
    if err := db.Delete(&takeAHistory).Error; err != nil {
        // ถ้ามีข้อผิดพลาดในการลบข้อมูล ให้ส่งข้อความ error กลับ
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete TakeAHistory", "details": err.Error()})
        return
    }

    // ส่งข้อความสำเร็จเมื่อลบข้อมูลเสร็จ
    c.JSON(http.StatusOK, gin.H{"message": "TakeAHistory deleted successfully"})
}
