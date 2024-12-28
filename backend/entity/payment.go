package entity

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	PatientID                  uint       `json:"patient_id"`
	EmployeeID                 uint       `json:"employee_id"`
	MedicalRecordsID           uint       `json:"medical_records_id"`
	PaymentDate                time.Time  `json:"payment_date"`
	PaymentAmount             float64    `json:"payment_amount"`
	StatusPaymentID           uint       `json:"status_payment_id"`
	PaymentServiceID          string     `json:"payment_service_id"`
	PaymentMethodID             uint     `json:"payment_method_id"`
	PaymentNotes              string     `json:"payment_notes"`
	PatientRoomID             uint       `json:"patient_room_id"` // เพิ่มฟิลด์เชื่อมโยง PatientRoom

	// Foreign Key relationships
	Patient                   Patient             `gorm:"foreignKey:PatientID"`
	Employee                  Employee            `gorm:"foreignKey:EmployeeID"`
	MedicalRecords           MedicalRecords      `gorm:"foreignKey:MedicalRecordsID"`
	StatusPayment            StatusPayment       `gorm:"foreignKey:StatusPaymentID"`
	PatientRoom              PatientRoom         `gorm:"foreignKey:PatientRoomID"`
	PaymentMethod			 PaymentMethod		  `gorm:"foreignKey:PaymentMethodID"`		 // เพิ่มการเชื่อมโยงกับ PatientRoom

	// Many-to-Many relationship with MedicalEntitlement
	MedicalEntitlements   []MedicalEntitlement  `gorm:"many2many:payment_medical_entitlement;" json:"medical_entitlements"`

}
