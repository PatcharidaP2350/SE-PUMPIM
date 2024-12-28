package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicalEquipment struct {
    gorm.Model
    Name              string    // ชื่ออุปกรณ์
    Category          string    // หมวดหมู่
    EquipmentModel             string    // รุ่น
    SerialNumber      string    // หมายเลขซีเรียล
    Manufacturer      string    // ผู้ผลิต
    CountryOfOrigin   string    // ประเทศที่ผลิต

    StockQuantity     int       // จำนวนในสต็อก
    ReorderLevel      int       // ระดับเตือนสั่งซื้อใหม่
    PricePerUnit      float32   // ราคาต่อหน่วย
    ImportDate        time.Time // วันที่นำเข้า
    ExpiryDate        time.Time // วันที่หมดอายุ (อาจเป็น nil)

    LastMaintenance   time.Time // วันที่ตรวจเช็คครั้งล่าสุด
    MaintenanceSchedule string   // รอบการบำรุงรักษา
    MaintenanceHistory string    // ประวัติการซ่อม

    SupplierID        uint       // รหัสผู้จัดจำหน่าย
    Supplier          Supplier   `gorm:"foreignKey:SupplierID"`

    UsageInstructions string     // คำแนะนำการใช้งาน
    Warnings          string     // คำเตือน
    Purpose           string     // วัตถุประสงค์
    Image             string     // ลิงก์รูปภาพ
    Standards         string     // มาตรฐาน
}