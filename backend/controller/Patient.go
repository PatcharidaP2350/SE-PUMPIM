package controller

import (
	"net/http"
	"time"
	"SE-B6527075/config"
	"SE-B6527075/entity"

	"github.com/gin-gonic/gin"
)


// CreatePatient - ฟังก์ชันสำหรับสร้างข้อมูล Patient
func CreatePatient(c *gin.Context) {
	var patient entity.Patient

	// Bind JSON จากคำขอไปยัง Entity `Patient`
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบ GenderID และ BloodGroupID
	db := config.DB()
	var gender entity.Gender
	var bloodGroup entity.BloodGroup

	if err := db.First(&gender, patient.GenderID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender_id"})
		return
	}

	if err := db.First(&bloodGroup, patient.BloodGroupID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blood_group_id"})
		return
	}

	// ตรวจสอบ Diseases (ถ้ามี)
	var diseases []entity.Disease
	if len(patient.Diseases) > 0 {
		diseaseIDs := []uint{}
		for _, disease := range patient.Diseases {
			diseaseIDs = append(diseaseIDs, disease.ID)
		}
		if err := db.Where("id IN ?", diseaseIDs).Find(&diseases).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid disease_ids"})
			return
		}
	}

	// อัปเดตความสัมพันธ์ Diseases
	patient.Diseases = diseases

	// กำหนดค่าฟิลด์อื่นๆ ถ้าจำเป็น เช่น DateOfBirth (ในกรณีที่ไม่ส่งมา)
	if patient.DateOfBirth.IsZero() {
		patient.DateOfBirth = time.Now() // กำหนดเป็นวันที่ปัจจุบันหากไม่ระบุ
	}

	// บันทึกข้อมูลลงฐานข้อมูล
	if err := db.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	// ตอบกลับข้อมูล Patient ที่สร้างสำเร็จ
	c.JSON(http.StatusCreated, gin.H{"data": patient})
}


// GetPatient - ฟังก์ชันสำหรับดึงข้อมูล Patient ตาม ID
func GetPatient(c *gin.Context) {
	var patient entity.Patient
	patientID := c.Param("id") // รับ ID จาก URL parameter

	// ดึงข้อมูล Patient จากฐานข้อมูล
	db := config.DB()
	if err := db.Preload("Gender").Preload("BloodGroup").Preload("Diseases").First(&patient, patientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// ส่งข้อมูล Patient กลับในรูปแบบ JSON
	c.JSON(http.StatusOK, gin.H{
		"id":            patient.ID,
		"nation_id":     patient.NationID,
		"first_name":    patient.FirstName,
		"last_name":     patient.LastName,
		"date_of_birth": patient.DateOfBirth,
		"address":       patient.Address,
		"phone_number":  patient.PhoneNumber,
		"gender":        patient.Gender.GenderName,     // แสดงชื่อ Gender
		"blood_group":   patient.BloodGroup.BloodGroup, // แสดงชื่อ Blood Group
		"diseases":      patient.Diseases,              // แสดงข้อมูล Disease
		"created_at":    patient.CreatedAt,
	})
}


// ListPatient - ฟังก์ชันสำหรับดึงข้อมูลรายชื่อผู้ป่วยทั้งหมด
func ListPatient(c *gin.Context) {
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ป่วยทั้งหมด
	var patients []entity.Patient

	// ดึงข้อมูลผู้ป่วยทั้งหมดจากฐานข้อมูล
	if err := config.DB().Preload("Gender").Preload("BloodGroup").Preload("Diseases").Find(&patients).Error; err != nil {
		// ถ้ามีข้อผิดพลาดในการดึงข้อมูล
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patients"})
		return
	}

	// สร้างตัวแปรสำหรับเก็บข้อมูลผู้ป่วยที่ต้องการส่งกลับ
	var patientList []gin.H

	// แปลงข้อมูลผู้ป่วยให้มีรูปแบบเหมาะสมสำหรับการส่งกลับ
	for _, patient := range patients {
		patientList = append(patientList, gin.H{
			"id":            patient.ID,
			"nation_id":     patient.NationID,
			"first_name":    patient.FirstName,
			"last_name":     patient.LastName,
			"date_of_birth": patient.DateOfBirth,
			"address":       patient.Address,
			"phone_number":  patient.PhoneNumber,
			"gender":        patient.Gender.GenderName,     // แสดงชื่อ Gender
			"blood_group":   patient.BloodGroup.BloodGroup, // แสดงชื่อ BloodGroup
			"diseases":      patient.Diseases,              // แสดงข้อมูล Disease
			"created_at":    patient.CreatedAt,
		})
	}

	// ส่งข้อมูลรายชื่อผู้ป่วยทั้งหมดกลับ
	c.JSON(http.StatusOK, gin.H{"data": patientList})
}


// DeletePatient - ฟังก์ชันสำหรับลบข้อมูลผู้ป่วยตาม ID
func DeletePatient(c *gin.Context) {
	// รับค่า ID ของผู้ป่วยจาก URL parameters
	patientID := c.Param("id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ป่วยที่ต้องการลบ
	var patient entity.Patient

	// ค้นหาผู้ป่วยตาม ID ที่รับมา
	if err := config.DB().First(&patient, patientID).Error; err != nil {
		// ถ้าหากไม่พบผู้ป่วยที่มี ID นี้
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// ลบข้อมูลผู้ป่วย
	if err := config.DB().Delete(&patient).Error; err != nil {
		// ถ้ามีข้อผิดพลาดในการลบข้อมูล
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}

	// ส่งข้อความตอบกลับว่าได้ลบข้อมูลสำเร็จ
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}


// UpdatePatient - ฟังก์ชันสำหรับอัปเดตข้อมูลบางฟิลด์ของผู้ป่วย
func UpdatePatient(c *gin.Context) {
	// รับค่า ID ของผู้ป่วยจาก URL parameters
	patientID := c.Param("id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ป่วยที่ต้องการอัปเดต
	var patient entity.Patient

	// ค้นหาผู้ป่วยตาม ID ที่รับมา
	if err := config.DB().First(&patient, patientID).Error; err != nil {
		// ถ้าหากไม่พบผู้ป่วยที่มี ID นี้
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Bind JSON จากคำขอไปยัง Entity `Patient` เพื่ออัปเดตข้อมูลบางฟิลด์
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบข้อมูลที่ได้รับมาเพียงบางฟิลด์ที่อัปเดต
	// ถ้าได้รับ GenderID หรือ BloodGroupID จะต้องตรวจสอบก่อน
	if patient.GenderID != 0 {
		var gender entity.Gender
		if err := config.DB().First(&gender, patient.GenderID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender_id"})
			return
		}
	}

	if patient.BloodGroupID != 0 {
		var bloodGroup entity.BloodGroup
		if err := config.DB().First(&bloodGroup, patient.BloodGroupID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blood_group_id"})
			return
		}
	}

	// ถ้ามีการส่ง diseases มาด้วย ให้ตรวจสอบและอัปเดต
	if len(patient.Diseases) > 0 {
		diseaseIDs := []uint{}
		for _, disease := range patient.Diseases {
			diseaseIDs = append(diseaseIDs, disease.ID)
		}
		var diseases []entity.Disease
		if err := config.DB().Where("id IN ?", diseaseIDs).Find(&diseases).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid disease_ids"})
			return
		}
		// อัปเดต diseases ของผู้ป่วย
		patient.Diseases = diseases
	}

	// อัปเดตเฉพาะฟิลด์ที่ได้รับการส่งมา
	if err := config.DB().Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	// ตอบกลับข้อมูล Patient ที่อัปเดตสำเร็จ
	c.JSON(http.StatusOK, gin.H{"data": patient})
}
