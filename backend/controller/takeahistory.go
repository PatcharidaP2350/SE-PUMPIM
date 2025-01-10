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
	fmt.Println("Creating or Updating TakeAHistory")
	db := config.DB()

	// กำหนดข้อมูลที่รับจาก Request
	var input struct {
		Weight                   float32    `json:"weight" binding:"required"`
		Height                    float32    `json:"height" binding:"required"`
		PreliminarySymptoms      string     `json:"preliminary_symptoms" binding:"required"`
		SystolicBloodPressure    uint       `json:"systolic_blood_pressure" binding:"required"`
		DiastolicBloodPressure   uint       `json:"diastolic_blood_pressure" binding:"required"`
		PulseRate                uint       `json:"pulse_rate" binding:"required"`
		Smoking                  bool     `json:"smoking"`
		DrinkAlcohol             bool     `json:"drink_alcohol"`
		LastMenstruationDate     time.Time  `json:"last_menstruation_date"`
		PatientID                uint       `json:"patient_id"`
		EmployeeID               uint       `json:"employee_id" binding:"required"`
		AppointmentID			 *uint     `json:"appointment_id"`
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
		Height:                   input.Height,
		PreliminarySymptoms:     input.PreliminarySymptoms,
		SystolicBloodPressure:   input.SystolicBloodPressure,
		DiastolicBloodPressure:  input.DiastolicBloodPressure,
		PulseRate:               input.PulseRate,
		Smoking:                 input.Smoking,
		DrinkAlcohol:            input.DrinkAlcohol,
		LastMenstruationDate:    input.LastMenstruationDate,
		
		MedicalRecordsID:       nil, // ปรับเปลี่ยนตามความเหมาะสม
		Date:                    time.Now(),
		PatientID:               input.PatientID,
		EmployeeID:              input.EmployeeID,
		AppointmentID:		     input.AppointmentID,
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


func ListTakeAHistory(c *gin.Context) {                   //----------------yes------------------//
    db := config.DB() // เรียกใช้การเชื่อมต่อฐานข้อมูลจาก config

    var takeAHistories []entity.TakeAHistory

    // ดึงข้อมูลทั้งหมดจากฐานข้อมูล พร้อม preload ความสัมพันธ์
    if err := db.
        Preload("Patient").
        Preload("Employee").
        Preload("MedicalRecords").
        Preload("Appointments").
        Find(&takeAHistories).Error; err != nil {
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
		Height                    *float32   `json:"height"`
		PreliminarySymptoms      *string    `json:"preliminary_symptoms"`
		SystolicBloodPressure    *uint      `json:"systolic_blood_pressure"`
		DiastolicBloodPressure   *uint      `json:"diastolic_blood_pressure"`
		PulseRate                *uint      `json:"pulse_rate"`
		Smoking                  *bool    `json:"smoking"`
		DrinkAlcohol             *bool     `json:"drink_alcohol"`
		QueueNumber    *string    `json:"queue_number"`  
		Date      *time.Time    `json:"date"`      
		QueueStatus         *string    `json:"queue_status"`
		LastMenstruationDate     *time.Time `json:"last_menstruation_date"`
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
	if input.Height != nil {
		take.Height = *input.Height
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

	if input.DrinkAlcohol != nil {
		take.DrinkAlcohol = *input.DrinkAlcohol
	}
	if input.LastMenstruationDate != nil {
		take.LastMenstruationDate = *input.LastMenstruationDate
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

func UpdatePatientDisease(c *gin.Context) {
	// รับ JSON จากคำขอ
	var input struct {
		PatientID uint   `json:"patient_id"` // ID ของผู้ป่วย
		DiseaseID    []uint `json:"disease_id"`    // ID ของโรคที่ต้องการอัปเดต
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "คำขอไม่ถูกต้อง ไม่สามารถแปลง payload ได้"})
		return
	}
  
	db := config.DB()

	// ตรวจสอบว่า DrugID ไม่เป็นค่าว่าง
	if len(input.DiseaseID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาระบุยาอย่างน้อยหนึ่งรายการ"})
		return
	}

	// Start a transaction to ensure atomicity
	tx := db.Begin()

	// ลบยาเก่าทั้งหมดสำหรับผู้ป่วย
	if err := tx.Exec("DELETE FROM patient_diseases WHERE patient_id = ?", input.PatientID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถลบข้อมูลโรคได้"})
		return
	}

	// เพิ่มยาใหม่ทั้งหมดสำหรับผู้ป่วย
	for _, diseaseID := range input.DiseaseID {
		if err := tx.Exec("INSERT INTO patient_diseases (patient_id, disease_id) VALUES (?, ?)", input.PatientID, diseaseID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถเพิ่มโรคใหม่ได้"})
			return
		}
	}

	// Commit the transaction
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "อัปเดตสำเร็จ"})
}