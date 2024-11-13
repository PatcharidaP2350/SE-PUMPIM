package entity

import (
    "time"
    "gorm.io/gorm"
)

// Entity สำหรับบันทึกการรักษา (MedicalRecord)
type MedicalRecord struct {
    gorm.Model
    Diagnosis     string    // วินิจฉัยโรค
    Treatment     string    // การรักษาที่ให้
	CurrentSymptoms string  // อาการในปัจจุบันที่ผู้ป่วยแจ้ง
	Allergies      string   // อาการแพ้ที่ผู้ป่วยมี
	Instructions    string  //คำแนะนำ
    VisitDate     time.Time // วันที่บันทึกการรักษา
    EmployeeID    uint      // FK ไปยัง Employee (แพทย์หรือเจ้าหน้าที่ที่ทำการบันทึก)
    PatientID     uint      // FK ไปยัง Patient (ผู้ป่วยที่ได้รับการรักษา)
}