package controller

import (
	"SE-B6527075/config"
	"SE-B6527075/entity"
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	
)

func ListTekeAHistoryformedicalrecord(c *gin.Context) {
	db := config.DB()
	var tekeahistory []struct {
		ID                     uint 
		Weight                 float32
		Height                  float32
		PreliminarySymptoms    string
		SystolicBloodPressure  uint
		DiastolicBloodPressure uint
		PulseRate              uint
		LastMenstruationDate   time.Time `json:"-"`
		FormattedLastMenstruationDate string `json:"FormattedLastMenstruationDate"`
		Date                   time.Time `json:"-"`
		FormattedDate		string  `json:"Date"`
		FirstName              string
		LastName               string
		DateOfBirth            time.Time `json:"-"`
		GenderName             string
		// Count 					uint
		Age					   int 				//ต้องใช้เป็นintเท่านั้น
	}
	result := db.Model(&entity.TakeAHistory{}).
		Select("take_a_histories.id", "patients.first_name", "patients.last_name", "genders.gender_name", "patients.date_of_birth", "take_a_histories.weight", "take_a_histories.height", "take_a_histories.preliminary_symptoms", "take_a_histories.diastolic_blood_pressure", "take_a_histories.systolic_blood_pressure", "take_a_histories.pulse_rate", "take_a_histories.date").
		Joins("inner join patients on take_a_histories.patient_id = patients.id ").
		Joins("inner join genders on patients.gender_id = genders.id").
		Where("take_a_histories.medical_records_id IS NULL and take_a_histories.appointment_id IS NULL").Find(&tekeahistory)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	for i := range tekeahistory {
		tekeahistory[i].Age = calculateAge(tekeahistory[i].DateOfBirth)
		tekeahistory[i].FormattedDate = tekeahistory[i].Date.Format("02-01-2006 15:04:05")
		tekeahistory[i].FormattedLastMenstruationDate = tekeahistory[i].LastMenstruationDate.Format("02-01-2006")
	}

	c.JSON(http.StatusOK, &tekeahistory)

}

func ListSavemedicalrecord(c *gin.Context) {
	db := config.DB()
	var tekeahistory []struct {
		ID                     uint 			`json:"ID"`
		Weight                 float32
		Height                  float32
		PreliminarySymptoms    string
		SystolicBloodPressure  uint
		DiastolicBloodPressure uint
		PulseRate              uint
		LastMenstruationDate   time.Time 		`json:"-"`
		FormattedLastMenstruationDate string    `json:"FormattedLastMenstruationDate"`
		Date                   time.Time 		`json:"-"`
		FormattedDate		   string  			`json:"Date"`
		FirstName              string
		LastName               string
		DateOfBirth            time.Time 		`json:"-"`
		GenderName             string
		SeverityName  		   string
		EfirstName             string            
		ElastName              string            
		Age					   int 				//ต้องใช้เป็นintเท่านั้น
		MID                    int 				`json:"MID"`
		PatientPicture			string
	}
	result := db.Model(&entity.TakeAHistory{}).
		Select("take_a_histories.id","patients.patient_picture","medical_records.id as m_id","patients.first_name", "patients.last_name", "genders.gender_name", "patients.date_of_birth", "take_a_histories.weight", "take_a_histories.height", "take_a_histories.preliminary_symptoms", "take_a_histories.diastolic_blood_pressure", "take_a_histories.systolic_blood_pressure", "take_a_histories.pulse_rate", "medical_records.date","employees.first_name as efirst_name","employees.last_name as elast_name","severities.severity_name","take_a_histories.last_menstruation_date").
		Joins("inner join patients on take_a_histories.patient_id = patients.id ").
		Joins("inner join genders on patients.gender_id = genders.id").
		Joins("inner join medical_records on take_a_histories.medical_records_id = medical_records.id").
		Joins("inner join severities on medical_records.severity_id = severities.id").
		Joins("inner join employees on medical_records.employee_id = employees.id").
		Where("take_a_histories.medical_records_id IS NOT NULL and medical_records.deleted_at IS NULL").Find(&tekeahistory)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	for i := range tekeahistory {
		tekeahistory[i].Age = calculateAge(tekeahistory[i].DateOfBirth)
		tekeahistory[i].FormattedDate = tekeahistory[i].Date.Format("02-01-2006")
// จัดการฟิลด์ LastMenstruationDate
if !tekeahistory[i].LastMenstruationDate.IsZero() {
	tekeahistory[i].FormattedLastMenstruationDate = tekeahistory[i].LastMenstruationDate.Format("02-01-2006")
} else {
	tekeahistory[i].FormattedLastMenstruationDate = "ไม่มีข้อมูล"
}

	}
	

	c.JSON(http.StatusOK, &tekeahistory)

}

func calculateAge(DateOfBirth time.Time) int {
	now := time.Now()
	years := now.Year() - DateOfBirth.Year()

	// ตรวจสอบว่าเดือนและวันเกิดของปีนี้มาถึงหรือยัง ถ้ายังก็ลดอายุลง 1
	if now.YearDay() < DateOfBirth.YearDay() {
		years--
	}
	return years
}

func GetTakeAHistoryForMedicalRecordByID(c *gin.Context) {
	ID := c.Param("id")
	db := config.DB()

	// โครงสร้างสำหรับจัดเก็บข้อมูลผลลัพธ์
	var takeAHistory []struct {
		ID                            uint      `json:"ID"`
		Weight                        float32   `json:"Weight"`
		Height                         float32   `json:"Height"`
		PreliminarySymptoms           string    `json:"PreliminarySymptoms"`
		SystolicBloodPressure         uint      `json:"SystolicBloodPressure"`
		DiastolicBloodPressure        uint      `json:"DiastolicBloodPressure"`
		PulseRate                     uint      `json:"PulseRate"`
		LastMenstruationDate          time.Time `json:"LastMenstruationDate"`
		FormattedLastMenstruationDate string    `json:"FormattedLastMenstruationDate"`
		Date                          time.Time `json:"Date"`
		FirstName                     string    `json:"FirstName"`
		LastName                      string    `json:"LastName"`
		DateOfBirth                   time.Time `json:"DateOfBirth"`
		GenderName                    string    `json:"GenderName"`
		Smoking                       string    `json:"Smoking"`
		DrinkAlcohol                  string    `json:"DrinkAlcohol"`
		BloodGroup                    string    `json:"BloodGroup"`
		DiseaseNames                  string    `json:"DiseaseNames"`
		Age                           int       `json:"Age"`
		DrugNames                     string    `json:"DrugNames"`
		PID							  uint  	`json:"PID"`
		Diagnose					string	`json:"Diagnose"`
		AID					string	`json:"AID"`
		PatientPicture		string `json:"PatientPicture"`
	}

	// คำสั่ง SQL ที่แก้ไขให้ใช้ DISTINCT ใน GROUP_CONCAT
	result := db.Model(&entity.TakeAHistory{}).
		Select(`take_a_histories.id,
				patients.patient_picture,
				patients.id AS p_id,
				patients.first_name,
				patients.last_name,
				genders.gender_name,
				patients.date_of_birth,
				take_a_histories.weight,
				take_a_histories.height,
				take_a_histories.preliminary_symptoms,
				take_a_histories.diastolic_blood_pressure,
				take_a_histories.systolic_blood_pressure,
				take_a_histories.last_menstruation_date,
				take_a_histories.pulse_rate,
				take_a_histories.date,
				take_a_histories.smoking,
				take_a_histories.drink_alcohol,
				blood_groups.blood_group,
				medical_records.diagnose,
				take_a_histories.appointment_id as a_id,
				GROUP_CONCAT(DISTINCT diseases.disease_name) AS disease_names,
				GROUP_CONCAT(DISTINCT drugs.drug_name) AS drug_names`).
		Joins("LEFT JOIN patients ON take_a_histories.patient_id = patients.id").
		Joins("LEFT JOIN genders ON patients.gender_id = genders.id").
		Joins("LEFT JOIN blood_groups ON patients.blood_group_id = blood_groups.id").
		Joins("LEFT JOIN patient_diseases ON patient_diseases.patient_id = patients.id").
		Joins("LEFT JOIN diseases ON patient_diseases.disease_id = diseases.id").
		Joins("LEFT JOIN patient_drug ON patient_drug.patient_id = patients.id").
		Joins("LEFT JOIN drugs ON patient_drug.drug_id = drugs.id").
		Joins("LEFT join appointments on take_a_histories.appointment_id = appointments.id").
		Joins("LEFT join medical_records on appointments.medical_records_id = medical_records.id").
		Where("take_a_histories.id = ?", ID).
		Group("take_a_histories.id").
		Scan(&takeAHistory)

	// จัดการกรณีที่เกิดข้อผิดพลาด
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// การจัดรูปแบบข้อมูลเพิ่มเติม
	for i := range takeAHistory {
		// คำนวณอายุ
		takeAHistory[i].Age = calculateAge(takeAHistory[i].DateOfBirth)

		// จัดการฟิลด์ LastMenstruationDate
		if !takeAHistory[i].LastMenstruationDate.IsZero() {
			takeAHistory[i].FormattedLastMenstruationDate = takeAHistory[i].LastMenstruationDate.Format("02-01-2006")
		} else {
			takeAHistory[i].FormattedLastMenstruationDate = "ไม่มีข้อมูล"
		}

		// // จัดการฟิลด์ DiseaseNames และ DrugNames
		// if takeAHistory[i].DiseaseNames == "" {
		// 	takeAHistory[i].DiseaseNames = "ไม่มีข้อมูลโรค"
		// }
		// if takeAHistory[i].DrugNames == "" {
		// 	takeAHistory[i].DrugNames = "ไม่มีข้อมูลยา"
		// }
	}

	// ส่งข้อมูลกลับในรูปแบบ JSON
	c.JSON(http.StatusOK, takeAHistory)
}


func CreateMedicalRecord(c *gin.Context) {
	fmt.Println("Creating or Updating Medical Record")

	// ดึง ID จาก URL parameter
	idParam := c.Param("id")

	var input struct {
		SymptomsDetails string  `json:"symptoms_details"` // อาการที่พบ
		CheckResults    string  `json:"check_results"`    // ผลการตรวจ
		Diagnose        string  `json:"diagnose"`         // การวินิจฉัย
		OrderMedicine   string  `json:"order_medicine"`   // ยาที่สั่ง
		Instructions    string  `json:"instructions"`     // คำแนะนำ
		Price           float64 `json:"price"`            // ราคา
		Severity     	uint    `json:"severity"`      	  // ระดับความรุนแรง
		Schedule        string    `json:"schedule"`      // ID ของตารางนัด
		EmployeeID      uint    `json:"employee_id"`      // ID ของพนักงาน
		Diseases        []uint  `json:"disease"`         // รายการ ID ของโรค
		Image           []string `json:"image"`           // Base64 รูปภาพ
		Hospitalization      string    `json:"hospitalization"`      // ID ของตารางนัด
	}
	// ดึงข้อมูล JSON จากคำขอ (Request) และตรวจสอบความถูกต้อง
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ดึงข้อมูล SeverityID จากตาราง Severity
	var severity  entity.Severity
	db := config.DB()
	if err := db.Where("severity_level = ?", input.Severity).First(&severity).Error; err != nil {
		fmt.Println("Error fetching severity:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid severity level"})
		return
	}
	var schedule  entity.Schedule
	if err := db.Where("schedule_name = ?", input.Schedule).First(&schedule).Error; err != nil {
		fmt.Println("Error fetching schedule:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule"})
		return
	}

	var hospitalization  entity.Hospitalization
	if err := db.Where("hospitalization_name = ?", input.Hospitalization).First(&hospitalization).Error; err != nil {
		fmt.Println("Error fetching hospitalization:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hospitalization"})
		return
	}
	// สร้างข้อมูล MedicalRecord ใหม่
	medicalrecord := entity.MedicalRecords{
		SymptomsDetails: input.SymptomsDetails,
		CheckResults:    input.CheckResults,
		Diagnose:        input.Diagnose,
		OrderMedicine:   input.OrderMedicine,
		Instructions:    input.Instructions,
		Date:            time.Now(), // ใช้วันที่ปัจจุบัน
		Price:           input.Price,
		SeverityID:      severity.ID,
		ScheduleID:      schedule.ID,
		EmployeeID:      input.EmployeeID,
		HospitalizationID: hospitalization.ID,
	}
	// ดึงข้อมูลโรค (Diseases) ที่สัมพันธ์กับ ID ที่ผู้ใช้ส่งมา
	var diseases []entity.Disease
	if len(input.Diseases) > 0 {
		if err := db.Where("id IN ?", input.Diseases).Find(&diseases).Error; err != nil {
			fmt.Println("Error fetching diseases:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid disease IDs"})
			return
		}
	}
	medicalrecord.Diseases = diseases

	// บันทึก MedicalRecord ลงฐานข้อมูล
	if err := db.Create(&medicalrecord).Error; err != nil {
		fmt.Println("Error saving medical record:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save medical record"})
		return
	}

	// ใช้ ID ที่สร้างขึ้นสำหรับ MedicalRecord
	newRecordID := medicalrecord.ID

	// ถ้ามีการส่ง ID ใน URL มาด้วย (สำหรับอัปเดตตาราง TakeAHistory)
	if idParam != "" {
		// ค้นหา TakeAHistory ที่มี id ตรงกับ idParam
		var history entity.TakeAHistory
		if err := db.Where("id = ?", idParam).First(&history).Error; err != nil {
			fmt.Println("Error finding TakeAHistory:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TakeAHistory ID"})
			return
		}

		// อัปเดต MedicalRecordsID ใน TakeAHistory ที่มี ID ตรงกับ idParam
		if err := db.Model(&history).Update("medical_records_id", newRecordID).Error; err != nil {
			fmt.Println("Error updating TakeAHistory:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update TakeAHistory"})
			return
		}
		fmt.Println("TakeAHistory updated successfully")
	} else {
		// หากไม่มี ID ใน URL ให้สร้างเรคคอร์ดใหม่ในตาราง TakeAHistory
		otherTable := entity.TakeAHistory{
			MedicalRecordsID: &newRecordID,
		}
		if err := db.Create(&otherTable).Error; err != nil {
			fmt.Println("Error creating TakeAHistory:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create TakeAHistory"})
			return
		}
		fmt.Println("TakeAHistory created successfully")
	}

	// บันทึกรูปภาพในรูปแบบ Base64
if len(input.Image) > 0 {
    var image []entity.MedicalImage

    for _, imgBase64 := range input.Image {
        // ตรวจสอบ Base64 เบื้องต้น
        if imgBase64 == "" {
            fmt.Println("Empty Base64 string found, skipping...")
            continue
        }

        image = append(image, entity.MedicalImage{
            MedicalRecordsID: newRecordID,
            Image:            imgBase64,
        })
    }

    // บันทึกรูปภาพทั้งหมดในครั้งเดียว
    if len(image) > 0 {
        if err := db.Create(&image).Error; err != nil {
            fmt.Println("Error saving images:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save images"})
            return
        }
        fmt.Println("Images saved successfully")
    } else {
        fmt.Println("No valid images to save")
    }
} else {
    fmt.Println("No images provided")
}


	// ตอบกลับสำเร็จ
	c.JSON(http.StatusCreated, gin.H{
		"message": "Medical record created and TakeAHistory handled successfully",
		"data":    medicalrecord,
	})
}

// func CreateTakeARecord(c *gin.Context) {
// 	fmt.Println("Creating or Updating Medical Record")
// 	db := config.DB()

// 	// กำหนดข้อมูลที่รับจาก Request
// 	var input struct {
// 		Weight                   float32    `json:"weight" binding:"required"`
// 		Height                    float32    `json:"height" binding:"required"`
// 		PreliminarySymptoms      string     `json:"preliminary_symptoms" binding:"required"`
// 		SystolicBloodPressure    uint       `json:"systolic_blood_pressure" binding:"required"`
// 		DiastolicBloodPressure   uint       `json:"diastolic_blood_pressure" binding:"required"`
// 		PulseRate                uint       `json:"pulse_rate" binding:"required"`
// 		Smoking                  string     `json:"smoking" binding:"required"`
// 		LastMenstruationDate     time.Time  `json:"last_menstruation_date" binding:"required"`
// 		DrinkAlcohol             string     `json:"drink_alcohol" binding:"required"`
// 		PatientID                uint       `json:"patient_id" binding:"required"`
// 		EmployeeID               uint       `json:"employee_id" binding:"required"`
// 	}

// 	// ดึงข้อมูล JSON จากคำขอ (Request) และตรวจสอบความถูกต้อง
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		fmt.Println("Error binding JSON:", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
// 		return
// 	}

// 	// สร้าง record ใหม่จากข้อมูลที่ได้รับ
// 	take := entity.TakeAHistory{
// 		Weight:                  input.Weight,
// 		Height:                   input.Height,
// 		PreliminarySymptoms:     input.PreliminarySymptoms,
// 		SystolicBloodPressure:   input.SystolicBloodPressure,
// 		DiastolicBloodPressure:  input.DiastolicBloodPressure,
// 		PulseRate:               input.PulseRate,
// 		Smoking:                 input.Smoking,
// 		LastMenstruationDate:    input.LastMenstruationDate,
// 		DrinkAlcohol:            input.DrinkAlcohol,
// 		Date:                    time.Now(),
// 		MedicalRecordsID:       nil, // ปรับเปลี่ยนตามความเหมาะสม
// 		PatientID:               input.PatientID,
// 		EmployeeID:              input.EmployeeID,
// 	}

// 	// เริ่มต้นการเชื่อมต่อฐานข้อมูล
// 	if err := db.Create(&take).Error; err != nil {
// 		fmt.Println("Error saving medical record:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save medical record", "details": err.Error()})
// 		return
// 	}

// 	// ส่งการตอบกลับที่ประสบความสำเร็จ
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Medical record created and TakeAHistory handled successfully",
// 		"data":    take,
// 	})
// }

// func GetTekeAHistoryformedicalrecordByID(c *gin.Context) {
// 	ID := c.Param("id")
// 	var tekeahistorybyid entity.TakeAHistory

// 	db := config.DB()
// 	result := db.Preload("Patient").                          // โหลดความสัมพันธ์ Patient
// 		Preload("Patient.Diseases").  
// 		First(&tekeahistorybyid, ID)

// 	if result.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "TakeAHistory not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, tekeahistorybyid)
// }

// func CreateMedicalRecord(c *gin.Context) {
// 	fmt.Println("Creating Medical Record")

// 	var input struct {
// 		SymptomsDetails string    `json:"symptoms_details"` // อาการที่พบ
// 		CheckResults    string    `json:"check_results"`    // ผลการตรวจ
// 		Diagnose        string    `json:"diagnose"`         // การวินิจฉัย
// 		OrderMedicine   string    `json:"order_medicine"`   // ยาที่สั่ง
// 		Instructions    string    `json:"instructions"`    // คำแนะนำ
// 		Price           float64   `json:"price"`           // ราคา
// 		SeverityID      uint      `json:"severity_id"`     // ระดับความรุนแรง
// 		ScheduleID      uint      `json:"schedule_id"`     // ID ของตารางนัด
// 		EmployeeID      uint      `json:"employee_id"`     // ID ของพนักงาน
// 		Diseases        []uint    `json:"diseases"`        // รายการ ID ของโรค
// 	}

// 	// ดึงข้อมูล JSON จากคำขอ (Request) และตรวจสอบความถูกต้อง
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		fmt.Println("Error binding JSON:", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// สร้างข้อมูล MedicalRecord ใหม่
// 	medicalrecord := entity.MedicalRecords{
// 		SymptomsDetails: input.SymptomsDetails,
// 		CheckResults:    input.CheckResults,
// 		Diagnose:        input.Diagnose,
// 		OrderMedicine:   input.OrderMedicine,
// 		Instructions:    input.Instructions,
// 		Date:            time.Now(), // ใช้วันที่ปัจจุบัน
// 		Price:           input.Price,
// 		SeverityID:      input.SeverityID,
// 		ScheduleID:      input.ScheduleID,
// 		EmployeeID:      input.EmployeeID,
// 	}

// 	// ดึงข้อมูลโรค (Diseases) ที่สัมพันธ์กับ ID ที่ผู้ใช้ส่งมา
// 	var diseases []entity.Disease
// 	db := config.DB()
// 	if len(input.Diseases) > 0 {
// 		if err := db.Where("id IN ?", input.Diseases).Find(&diseases).Error; err != nil {
// 			fmt.Println("Error fetching diseases:", err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid disease IDs"})
// 			return
// 		}
// 	}
// 	medicalrecord.Diseases = diseases

// 	// บันทึก MedicalRecord ลงฐานข้อมูล
// 	if err := db.Create(&medicalrecord).Error; err != nil {
// 		fmt.Println("Error saving medical record:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save medical record"})
// 		return
// 	}

// 	// ใช้ ID ที่สร้างขึ้นสำหรับ MedicalRecord อัปเดตตารางอื่น
// 	newRecordID := medicalrecord.ID
// 	otherTable := entity.TakeAHistory{
// 		MedicalRecordsID: &newRecordID, // ใช้ ID ที่เพิ่งสร้าง
// 	}

// 	// บันทึกข้อมูลลงในตารางอื่น
// 	if err := db.Create(&otherTable).Error; err != nil {
// 		fmt.Println("Error updating other table:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update other table"})
// 		return
// 	}

// 	// ตอบกลับสำเร็จ
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Medical record created and other table updated successfully",
// 		"data":    medicalrecord,
// 	})
// }

func DeleteTakeAHistoryForMedicalRecord(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM take_a_histories WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Take A Histories Successful"})
}

func DeleterMedicalRecord(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()

	// Update `deleted_at` field to mark as deleted (using current timestamp)
	if tx := db.Exec("UPDATE medical_records SET deleted_at = datetime('now') WHERE id = ? AND deleted_at IS NULL", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found or already deleted"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Soft Deleted Medical Record Successfully"})
}

func UpdateMedicalRecord(c *gin.Context) {
	MedicalRecordsID := c.Param("id")

	// โครงสร้างสำหรับรับข้อมูล JSON
	var update struct {
		SymptomsDetails string  `json:"symptoms_details"` // อาการที่พบ
		CheckResults    string  `json:"check_results"`    // ผลการตรวจ
		Diagnose        string  `json:"diagnose"`         // การวินิจฉัย
		OrderMedicine   string  `json:"order_medicine"`   // ยาที่สั่ง
		Instructions    string  `json:"instructions"`     // คำแนะนำ
		Price           float64 `json:"price"`            // ราคา
		SeverityID      uint    `json:"severity_id"`      // ระดับความรุนแรง
		ScheduleID      uint    `json:"schedule_id"`      // ID ของตารางนัด
		EmployeeID      uint    `json:"employee_id"`      // ID ของพนักงาน
		Diseases        []uint  `json:"diseases"`         // รายการ ID ของโรค
	}

	db := config.DB()

	// ตรวจสอบและแมปข้อมูล JSON ไปยัง Struct `update`
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลที่ส่งมาไม่ถูกต้อง"})
		return
	}

	// ค้นหาข้อมูล Medical Record ตาม ID
	var medicalrecord entity.MedicalRecords
	if err := db.First(&medicalrecord, MedicalRecordsID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลบันทึกการรักษา"})
		return
	}

	// อัปเดตฟิลด์ของ Medical Record
	medicalrecord.SymptomsDetails = update.SymptomsDetails
	medicalrecord.CheckResults = update.CheckResults
	medicalrecord.Diagnose = update.Diagnose
	medicalrecord.OrderMedicine = update.OrderMedicine
	medicalrecord.Instructions = update.Instructions
	medicalrecord.Price = update.Price
	medicalrecord.SeverityID = update.SeverityID
	medicalrecord.ScheduleID = update.ScheduleID
	medicalrecord.EmployeeID = update.EmployeeID
	medicalrecord.Date = time.Now() // อัปเดตวันที่

	// อัปเดตความสัมพันธ์กับ Diseases
	if update.Diseases != nil {
		var diseases []entity.Disease
		if len(update.Diseases) > 0 {
			// ตรวจสอบว่า Disease ที่ส่งมามีอยู่ในฐานข้อมูล
			if err := db.Where("id IN ?", update.Diseases).Find(&diseases).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบโรคบางรายการ"})
				return
			}
			// อัปเดตรายการ Diseases
			if err := db.Model(&medicalrecord).Association("Diseases").Replace(diseases); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถอัปเดตข้อมูลโรคได้"})
				return
			}
		} else {
			// หากไม่มี ID โรคให้ลบ Association ทั้งหมด
			if err := db.Model(&medicalrecord).Association("Diseases").Clear(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถลบข้อมูลโรคได้"})
				return
			}
		}
	}

	// บันทึกการอัปเดตลงฐานข้อมูล
	if err := db.Save(&medicalrecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการบันทึกข้อมูล"})
		return
	}

	// ส่งผลลัพธ์กลับ
	c.JSON(http.StatusOK, gin.H{"message": "อัปเดตข้อมูลสำเร็จ", "data": medicalrecord})
}

func GetSeverityMaxMin(c *gin.Context) {
	db := config.DB()
	var maxmin []struct {
		ID                  uint `json:"id"`
		MaxSeverityLevel		uint  `json:"max_severity_level"`
		MinSeverityLevel		uint  `json:"min_severity_level"`
		Range            uint `json:"range"`
	}
	result := db.Model(&entity.Severity{}).
		Select("id,max(severity_level) as max_severity_level,min(severity_level) as min_severity_level,((max(severity_level)-min(severity_level))/(count(severity_level)-1)) as range").Find(&maxmin)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &maxmin)

}

func GetSchedule(c *gin.Context) {
	db := config.DB()

	// Define a struct with two fields for the separate rows
	var schedule struct {
		ScheduleYes string `json:"schedule_yes"`
		ScheduleNo string `json:"schedule_no"`
	}

	// Retrieve the first row
	var scheduleRow1 struct {
		ScheduleName string
	}
	result1 := db.Model(&entity.Schedule{}).
		Select("schedule_name").Order("id").Limit(1).Find(&scheduleRow1)
	if result1.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result1.Error.Error()})
		return
	}

	// Retrieve the second row
	var scheduleRow2 struct {
		ScheduleName string
	}
	result2 := db.Model(&entity.Schedule{}).
		Select("schedule_name").Order("id").Offset(1).Limit(1).Find(&scheduleRow2)
	if result2.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result2.Error.Error()})
		return
	}

	// Assign the retrieved rows to the combined struct
	schedule.ScheduleYes = scheduleRow1.ScheduleName
	schedule.ScheduleNo = scheduleRow2.ScheduleName

	c.JSON(http.StatusOK, &schedule)
}

func UpdatePatientDrug(c *gin.Context) {
	// รับ JSON จากคำขอ
	var input struct {
		PatientID uint   `json:"patient_id"` // ID ของผู้ป่วย
		DrugID    []uint `json:"drug_id"`    // ID ของยาที่ต้องการอัปเดต
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "คำขอไม่ถูกต้อง ไม่สามารถแปลง payload ได้"})
		return
	}
  
	db := config.DB()

	// ตรวจสอบว่า DrugID ไม่เป็นค่าว่าง
	if len(input.DrugID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาระบุยาอย่างน้อยหนึ่งรายการ"})
		return
	}

	// Start a transaction to ensure atomicity
	tx := db.Begin()

	// ลบยาเก่าทั้งหมดสำหรับผู้ป่วย
	if err := tx.Exec("DELETE FROM patient_drug WHERE patient_id = ?", input.PatientID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถลบข้อมูลยาได้"})
		return
	}

	// เพิ่มยาใหม่ทั้งหมดสำหรับผู้ป่วย
	for _, drugID := range input.DrugID {
		if err := tx.Exec("INSERT INTO patient_drug (patient_id, drug_id) VALUES (?, ?)", input.PatientID, drugID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถเพิ่มยาใหม่ได้"})
			return
		}
	}

	// Commit the transaction
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "อัปเดตสำเร็จ"})
}

func ListTekeAHistoryforSchedule(c *gin.Context) {
	var input struct {
		EmployeeID uint `json:"employee_id"` // ID ของพนักงาน
	}

	// Bind the request body to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initialize the DB connection
	db := config.DB()

	// Define the struct for holding query results
	var tekeahistory []struct {
		ID                     uint
		Weight                 float32
		Height                  float32
		PreliminarySymptoms    string
		SystolicBloodPressure  uint
		DiastolicBloodPressure uint
		PulseRate              uint
		LastMenstruationDate   time.Time `json:"-"` // Don't return in the JSON response
		FormattedLastMenstruationDate string `json:"FormattedLastMenstruationDate"`
		Date                   time.Time `json:"-"`  // Don't return in the JSON response
		FormattedDate          string   `json:"Date"`
		FirstName              string
		LastName               string
		DateOfBirth            time.Time `json:"-"` // Don't return in the JSON response
		GenderName             string
		Age                    int
		EfirstName             string
		ElastName              string
		Diagnose				string
	}

	// Execute the SQL query
	result := db.Model(&entity.TakeAHistory{}).
		Select("take_a_histories.id", "employees.first_name as efirst_name", "employees.last_name as elast_name",
			"patients.first_name", "patients.last_name", "genders.gender_name", "patients.date_of_birth", 
			"take_a_histories.weight", "take_a_histories.height", "take_a_histories.preliminary_symptoms", 
			"take_a_histories.diastolic_blood_pressure", "take_a_histories.systolic_blood_pressure", 
			"take_a_histories.pulse_rate", "take_a_histories.date","medical_records.diagnose").
		Joins("inner join patients on take_a_histories.patient_id = patients.id").
		Joins("inner join genders on patients.gender_id = genders.id").
		Joins("inner join appointments on take_a_histories.appointment_id = appointments.id").
		Joins("inner join medical_records on appointments.medical_records_id = medical_records.id").
		Joins("inner join employees on appointments.employee_id = employees.id").
		Where("take_a_histories.medical_records_id IS NULL and appointments.employee_id = ?", input.EmployeeID).
		Find(&tekeahistory)

	// Handle potential errors
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Process the data
	for i := range tekeahistory {
		tekeahistory[i].Age = calculateAge(tekeahistory[i].DateOfBirth)
		tekeahistory[i].FormattedDate = tekeahistory[i].Date.Format("02-01-2006 15:04:05")
		if !tekeahistory[i].LastMenstruationDate.IsZero() {
			tekeahistory[i].FormattedLastMenstruationDate = tekeahistory[i].LastMenstruationDate.Format("02-01-2006")
		}
	}

	// Return the results as JSON
	c.JSON(http.StatusOK, &tekeahistory)
}

func GetHospitalization(c *gin.Context) {
	db := config.DB()

	// Define a struct with two fields for the separate rows
	var hospitalization struct {
		HospitalizationYes string `json:"hospitalization_yes"`
		HospitalizationNo string `json:"hospitalization_no"`
	}

	// Retrieve the first row
	var hospitalizationRow1 struct {
		HospitalizationName string
	}
	result1 := db.Model(&entity.Hospitalization{}).
		Select("hospitalization_name").Order("id").Limit(1).Find(&hospitalizationRow1)
	if result1.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result1.Error.Error()})
		return
	}

	// Retrieve the second row
	var hospitalizationRow2 struct {
		HospitalizationName string
	}
	result2 := db.Model(&entity.Hospitalization{}).
		Select("hospitalization_name").Order("id").Offset(1).Limit(1).Find(&hospitalizationRow2)
	if result2.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result2.Error.Error()})
		return
	}

	// Assign the retrieved rows to the combined struct
	hospitalization.HospitalizationYes = hospitalizationRow1.HospitalizationName
	hospitalization.HospitalizationNo = hospitalizationRow2.HospitalizationName

	c.JSON(http.StatusOK, &hospitalization)
}

func GetDetailMedicalRecordByID(c *gin.Context) {
	ID := c.Param("id")
	db := config.DB()

	// โครงสร้างสำหรับจัดเก็บข้อมูลผลลัพธ์
	var takeAHistory []struct {
		ID                            uint      `json:"ID"`
		Weight                        float32   `json:"Weight"`
		Height                         float32   `json:"Height"`
		PreliminarySymptoms           string    `json:"PreliminarySymptoms"`
		SystolicBloodPressure         uint      `json:"SystolicBloodPressure"`
		DiastolicBloodPressure        uint      `json:"DiastolicBloodPressure"`
		PulseRate                     uint      `json:"PulseRate"`
		LastMenstruationDate          time.Time `json:"LastMenstruationDate"`
		FormattedLastMenstruationDate string    `json:"FormattedLastMenstruationDate"`
		//Date                          time.Time `json:"Date"`
		FirstName                     string    `json:"FirstName"`
		LastName                      string    `json:"LastName"`
		DateOfBirth                   time.Time `json:"DateOfBirth"`
		GenderName                    string    `json:"GenderName"`
		Smoking                       string    `json:"Smoking"`
		DrinkAlcohol                  string    `json:"DrinkAlcohol"`
		BloodGroup                    string    `json:"BloodGroup"`
		DiseaseNames                  string    `json:"DiseaseNames"`
		Age                           int       `json:"Age"`
		DrugNames                     string    `json:"DrugNames"`
		PID							  uint  	`json:"PID"`
		Diagnose					  string	`json:"Diagnose"`
		//AID							  string	`json:"AID"`
		PatientPicture				  string 	`json:"PatientPicture"`
		SymptomsDetails				  string 	`json:"SymptomsDetails"`
		CheckResults				  string	`json:"CheckResults"`
		OrderMedicine				  string	`json:"OrderMedicine"`
		Instructions				  string	`json:"Instructions"`
		MDate						  time.Time	`json:"-"`
		FormattedMDate					string`json:"FormattedMDate"`
		Price						  float64	`json:"Price"`
		SeverityName				  string	`json:"SeverityName"`
		ScheduleName				  string	`json:"ScheduleName"`
		HospitalizationName				string `json:"HospitalizationName"`
		Image						  string	`json:"Image"`
		EFirstName                     string    `json:"EFirstName"`
		ELastName                      string    `json:"ELastName"`
		MedicalDiseaseNames                  string    `json:"MedicalDiseaseNames"`
	}

	// คำสั่ง SQL ที่แก้ไขให้ใช้ DISTINCT ใน GROUP_CONCAT
	result := db.Model(&entity.TakeAHistory{}).
		Select(`take_a_histories.id,
				patients.patient_picture,
				patients.id AS p_id,
				patients.first_name,
				patients.last_name,
				genders.gender_name,
				patients.date_of_birth,
				take_a_histories.weight,
				take_a_histories.height,
				take_a_histories.preliminary_symptoms,
				take_a_histories.diastolic_blood_pressure,
				take_a_histories.systolic_blood_pressure,
				take_a_histories.last_menstruation_date,
				take_a_histories.pulse_rate,
				take_a_histories.date,
				take_a_histories.smoking,
				take_a_histories.drink_alcohol,
				blood_groups.blood_group,
				medical_records.symptoms_details,
				medical_records.check_results,
				medical_records.diagnose,
				medical_records.order_medicine,
				medical_records.instructions,
				medical_records.date as m_date,
				medical_records.price,
				employees.first_name as e_first_name,
				employees.last_name as e_last_name,
				severities.severity_name,
				schedules.schedule_name,
				hospitalizations.hospitalization_name,
				take_a_histories.appointment_id as a_id,
				GROUP_CONCAT(DISTINCT medicaldiseases.disease_name) AS medical_disease_names,
				GROUP_CONCAT(DISTINCT medical_images.image) AS image,
				GROUP_CONCAT(DISTINCT diseases.disease_name) AS disease_names,
				GROUP_CONCAT(DISTINCT drugs.drug_name) AS drug_names`).
		Joins("LEFT JOIN patients ON take_a_histories.patient_id = patients.id").
		Joins("LEFT JOIN genders ON patients.gender_id = genders.id").
		Joins("LEFT JOIN blood_groups ON patients.blood_group_id = blood_groups.id").
		Joins("LEFT JOIN patient_diseases ON patient_diseases.patient_id = patients.id").
		Joins("LEFT JOIN diseases ON patient_diseases.disease_id = diseases.id").
		Joins("LEFT JOIN patient_drug ON patient_drug.patient_id = patients.id").
		Joins("LEFT JOIN drugs ON patient_drug.drug_id = drugs.id").
		Joins("LEFT join appointments on take_a_histories.appointment_id = appointments.id").
		Joins("LEFT join medical_records on take_a_histories.medical_records_id = medical_records.id").
		Joins("LEFT join severities on medical_records.severity_id = severities.id").
		Joins("LEFT join schedules on medical_records.schedule_id = schedules.id").
		Joins("LEFT join hospitalizations on medical_records.hospitalization_id = hospitalizations.id").
		Joins("LEFT join medical_images on medical_images.medical_records_id = medical_records.id").
		Joins("LEFT join employees on medical_records.employee_id = employees.id").
		Joins("LEFT JOIN medicalrecords_diseases ON medicalrecords_diseases.medical_records_id = medical_records.id").
		Joins("LEFT JOIN diseases as medicaldiseases ON medicalrecords_diseases.disease_id = medicaldiseases.id").
		Where("take_a_histories.id = ?", ID).
		Group("take_a_histories.id").
		Scan(&takeAHistory)

	// จัดการกรณีที่เกิดข้อผิดพลาด
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// การจัดรูปแบบข้อมูลเพิ่มเติม
	for i := range takeAHistory {
		// คำนวณอายุ
		takeAHistory[i].Age = calculateAge(takeAHistory[i].DateOfBirth)
		takeAHistory[i].FormattedMDate = takeAHistory[i].MDate.Format("02-01-2006 15:04:05")

		// จัดการฟิลด์ LastMenstruationDate
		if !takeAHistory[i].LastMenstruationDate.IsZero() {
			takeAHistory[i].FormattedLastMenstruationDate = takeAHistory[i].LastMenstruationDate.Format("02-01-2006")
		} else {
			takeAHistory[i].FormattedLastMenstruationDate = "ไม่มีข้อมูล"
		}

		// // จัดการฟิลด์ DiseaseNames และ DrugNames
		// if takeAHistory[i].DiseaseNames == "" {
		// 	takeAHistory[i].DiseaseNames = "ไม่มีข้อมูลโรค"
		// }
		// if takeAHistory[i].DrugNames == "" {
		// 	takeAHistory[i].DrugNames = "ไม่มีข้อมูลยา"
		// }
	}

	// ส่งข้อมูลกลับในรูปแบบ JSON
	c.JSON(http.StatusOK, takeAHistory)
}