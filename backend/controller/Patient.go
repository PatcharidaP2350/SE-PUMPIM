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

    db := config.DB()

    // ตรวจสอบ GenderID
    var gender entity.Gender
    if err := db.First(&gender, patient.GenderID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender_id"})
        return
    }

    // ตรวจสอบ BloodGroupID
    var bloodGroup entity.BloodGroup
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

    // ตรวจสอบฟิลด์ PatientPicture
    // ถ้าไม่ส่งค่ามา อนุญาตให้เป็น NULL ได้
    if patient.PatientPicture == nil {
        patient.PatientPicture = nil
    }

    // ตรวจสอบ Drug (ถ้ามี)
    var drugs []entity.Drug
    if len(patient.Drug) > 0 {
        drugIDs := []uint{}
        for _, drug := range patient.Drug {
            drugIDs = append(drugIDs, drug.ID)
        }
        if err := db.Where("id IN ?", drugIDs).Find(&drugs).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid drug_ids"})
            return
        }
    }

    // อัปเดตความสัมพันธ์ Drug
    patient.Drug = drugs

    // ตรวจสอบฟิลด์อื่นๆ เช่น DateOfBirth (ในกรณีที่ไม่ส่งมา)
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

	// ดึงข้อมูล Patient พร้อม Preload ความสัมพันธ์
	db := config.DB()
	if err := db.Preload("Gender").
		Preload("BloodGroup").
		Preload("Diseases").
		Preload("Drug").
		First(&patient, patientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// จัดรูปแบบข้อมูลเพื่อตอบกลับ
	response := gin.H{
		"ID":              patient.ID,
		"NationID":       patient.NationID,
		"FirstName":      patient.FirstName,
		"LastName":       patient.LastName,
		"DateOfBirth":   patient.DateOfBirth.Format("2006-01-02"), // แปลงวันที่
		"Address":         patient.Address,
		"PhoneNumber":    patient.PhoneNumber,
		"Gender": gin.H{
			"gender_id":   patient.Gender.ID,
			"gender_name": patient.Gender.GenderName,
		},
		"BloodGroup": gin.H{
			"blood_group_id":   patient.BloodGroup.ID,
			"blood_group_name": patient.BloodGroup.BloodGroup,
		},
		"PatientPicture": patient.PatientPicture, // สามารถเป็น nil ได้
		"Diseases": func() []gin.H {
			var diseases []gin.H
			for _, disease := range patient.Diseases {
				diseases = append(diseases, gin.H{
					"id":   disease.ID,
					"name": disease.DiseaseName,
				})
			}
			return diseases
		}(),
		"Drug": func() []gin.H {
			var drugs []gin.H
			for _, drug := range patient.Drug {
				drugs = append(drugs, gin.H{
					"id":   drug.ID,
					"name": drug.DrugName,
				})
			}
			return drugs
		}(),
		"created_at": patient.CreatedAt,
	}

	// ส่งข้อมูลกลับ
	c.JSON(http.StatusOK, response)
}


func GetPatientbyNationID(c *gin.Context) {
	var patient entity.Patient
	nationID := c.Param("nation_id") // รับ nation_id จาก URL parameter

	// ดึงข้อมูล Patient พร้อม Preload ความสัมพันธ์
	db := config.DB()
	if err := db.Preload("Gender").
		Preload("BloodGroup").
		Preload("Diseases").
		Preload("Drug").
		Where("nation_id = ?", nationID).
		First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// จัดรูปแบบข้อมูลเพื่อตอบกลับ
	response := gin.H{
		"ID":              patient.ID,
		"NationID":       patient.NationID,
		"FirstName":      patient.FirstName,
		"LastName":       patient.LastName,
		"DateOfBirth":   patient.DateOfBirth.Format("2006-01-02"), // แปลงวันที่
		"Address":         patient.Address,
		"PhoneNumber":    patient.PhoneNumber,
		"Gender": gin.H{
			"gender_id":   patient.Gender.ID,
			"gender_name": patient.Gender.GenderName,
		},
		"BloodGroup": gin.H{
			"blood_group_id":   patient.BloodGroup.ID,
			"blood_group": patient.BloodGroup.BloodGroup,
		},
		"PatientPicture": patient.PatientPicture, // สามารถเป็น nil ได้
		"Diseases": func() []gin.H {
			var diseases []gin.H
			for _, disease := range patient.Diseases {
				diseases = append(diseases, gin.H{
					"id":   disease.ID,
					"name": disease.DiseaseName,
				})
			}
			return diseases
		}(),
		"Drugs": func() []gin.H {
			var drugs []gin.H
			for _, drug := range patient.Drug {
				drugs = append(drugs, gin.H{
					"id":   drug.ID,
					"name": drug.DrugName,
				})
			}
			return drugs
		}(),
		"created_at": patient.CreatedAt,
	}

	// ส่งข้อมูลกลับ
	c.JSON(http.StatusOK, response)
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
			"ID":            patient.ID,
			"NationID":     patient.NationID,
			"FirstName":    patient.FirstName,
			"LastName":     patient.LastName,
			"DateOfBirth": patient.DateOfBirth,
			"Address":       patient.Address,
			"PhoneNumber":  patient.PhoneNumber,
			"Gender": gin.H{
			"id":   patient.Gender.ID,
			"name": patient.Gender.GenderName,
			},
			"BloodGroup": gin.H{
				"id":   patient.BloodGroup.ID,
				"name": patient.BloodGroup.BloodGroup,
			},
			"PatientPicture": patient.PatientPicture, // สามารถเป็น nil ได้
			"Diseases": func() []gin.H {
				var diseases []gin.H
				for _, disease := range patient.Diseases {
					diseases = append(diseases, gin.H{
						"id":   disease.ID,
						"name": disease.DiseaseName,
					})
				}
				return diseases
			}(),
			"Drug": func() []gin.H {
				var drugs []gin.H
				for _, drug := range patient.Drug {
					drugs = append(drugs, gin.H{
						"id":   drug.ID,
						"name": drug.DrugName,
					})
				}
				return drugs
			}(),
			"created_at": patient.CreatedAt,
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
	db := config.DB()
	if err := db.Preload("Diseases").Preload("Drug").Preload("TakeAHistorys").First(&patient, patientID).Error; err != nil {
		// ถ้าหากไม่พบผู้ป่วยที่มี ID นี้
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// ลบความสัมพันธ์ในตาราง Many-to-Many ก่อน
	if len(patient.Diseases) > 0 {
		if err := db.Model(&patient).Association("Diseases").Clear(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear diseases relationship"})
			return
		}
	}

	if len(patient.Drug) > 0 {
		if err := db.Model(&patient).Association("Drug").Clear(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear drug relationship"})
			return
		}
	}

	// ลบประวัติการรักษาที่เกี่ยวข้อง
	if len(patient.TakeAHistorys) > 0 {
		if err := db.Where("patient_id = ?", patient.ID).Delete(&entity.TakeAHistory{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete related history records"})
			return
		}
	}

	// ลบข้อมูลผู้ป่วย
	if err := db.Delete(&patient).Error; err != nil {
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
	if err := config.DB().Preload("Diseases").First(&patient, patientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// สร้างตัวแปรเพื่อรับข้อมูลใหม่จากคำขอ
	var updateData struct {
		NationID       *string        `json:"nation_id"`
		FirstName      *string        `json:"first_name"`
		LastName       *string        `json:"last_name"`
		DateOfBirth    *time.Time     `json:"date_of_birth"`
		Address        *string        `json:"address"`
		PhoneNumber    *string        `json:"phone_number"`
		GenderID       *uint          `json:"gender_id"`
		BloodGroupID   *uint          `json:"blood_group_id"`
		PatientPicture *string        `json:"patient_picture"`
		Diseases       []entity.Disease `json:"diseases"`
	}

	// Bind JSON จากคำขอไปยังตัวแปร updateData
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตข้อมูลตามฟิลด์ที่ได้รับ
	if updateData.NationID != nil {
		patient.NationID = *updateData.NationID
	}
	if updateData.FirstName != nil {
		patient.FirstName = *updateData.FirstName
	}
	if updateData.LastName != nil {
		patient.LastName = *updateData.LastName
	}
	if updateData.DateOfBirth != nil {
		patient.DateOfBirth = *updateData.DateOfBirth
	}
	if updateData.Address != nil {
		patient.Address = *updateData.Address
	}
	if updateData.PhoneNumber != nil {
		patient.PhoneNumber = *updateData.PhoneNumber
	}
	if updateData.GenderID != nil {
		var gender entity.Gender
		if err := config.DB().First(&gender, *updateData.GenderID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender_id"})
			return
		}
		patient.GenderID = *updateData.GenderID
	}
	if updateData.BloodGroupID != nil {
		var bloodGroup entity.BloodGroup
		if err := config.DB().First(&bloodGroup, *updateData.BloodGroupID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blood_group_id"})
			return
		}
		patient.BloodGroupID = *updateData.BloodGroupID
	}
	if updateData.PatientPicture != nil {
		patient.PatientPicture = updateData.PatientPicture
	}

	// อัปเดต Diseases ถ้ามีการส่งมาด้วย
	if len(updateData.Diseases) > 0 {
		diseaseIDs := []uint{}
		for _, disease := range updateData.Diseases {
			diseaseIDs = append(diseaseIDs, disease.ID)
		}
		var diseases []entity.Disease
		if err := config.DB().Where("id IN ?", diseaseIDs).Find(&diseases).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid disease_ids"})
			return
		}
		// อัปเดต Diseases ของผู้ป่วย
		patient.Diseases = diseases
	}

	// บันทึกการเปลี่ยนแปลง
	if err := config.DB().Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	// ส่งข้อมูลผู้ป่วยที่อัปเดตกลับในรูปแบบ JSON
	c.JSON(http.StatusOK, gin.H{"data": patient})
}
