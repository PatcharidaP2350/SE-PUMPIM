package entity

import (
	"time"

	"gorm.io/gorm"
)

type Drug struct {
    gorm.Model
    DrugName       string    // ชื่อยา
    Category       string    // หมวดหมู่
    Formulation    string    // รูปแบบยา
    Dosage         string    // ขนาดยา
    RegistrationNo string    // หมายเลขทะเบียนยา

    StockQuantity  int       // จำนวนในสต็อก
    ReorderLevel   int       // ระดับเตือนสั่งซื้อใหม่
    PricePerUnit   float32   // ราคาต่อหน่วย
    ImportDate     time.Time // วันที่นำเข้า
    ExpiryDate     time.Time // วันหมดอายุ

    Manufacturer   string    // ผู้ผลิต
    CountryOfOrigin string   // ประเทศที่ผลิต
    BatchNumber    string    // ล็อตที่ผลิต

    SupplierID     uint       // รหัสผู้จัดจำหน่าย
    Supplier       Supplier   `gorm:"foreignKey:SupplierID"`

    UsageInstructions string  // คำแนะนำการใช้ยา
    Indications       string  // ข้อบ่งใช้
    Contraindications string  // ข้อห้ามใช้
    SideEffects       string  // ผลข้างเคียง

    DrugImage        string   // รูปภาพของยา
    Barcode          string   // รหัสบาร์โค้ด
    Patient   []Patient `gorm:"many2many:patient_drug;" json:"patient"` // Many-to-Many relationship with Pateint
}