package entity

import "gorm.io/gorm"

type MedicalEntitlement struct {
	gorm.Model
	MedicalEntitlementID      uint   `json:"medical_entitlement_id"`      // รหัสสิทธิการรักษาพยาบาล
	MedicalEntitlementName    string `json:"medical_entitlement_name"`    // ชื่อสิทธิการรักษาพยาบาล
	MedicalEntitlementType    string `json:"medical_entitlement_type"`    // ประเภทของสิทธิการรักษาพยาบาล
	MedicalEntitlementUsageLimit float64 `json:"medical_entitlement_usage_limit"` // ขีดจำกัดการใช้งานสิทธิ
	MedicalEntitlementProviderName string `json:"medical_entitlement_provider_name"` // ชื่อผู้ให้บริการสิทธิ
	MedicalEntitlementProviderContact string `json:"medical_entitlement_provider_contact"` // ช่องทางติดต่อผู้ให้บริการสิทธิ

	// Many-to-Many relationship with Payment
	Payments []Payment `gorm:"many2many:payment_medical_entitlement;" json:"payments"`
}
