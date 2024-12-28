import { SupplierInterface } from "./ISupplier";

export interface DrugsInterface {
    ID?: number;

    DrugName: string; // ชื่อยา
    Category: string; // หมวดหมู่
    Formulation: string; // รูปแบบยา
    Dosage: string; // ขนาดยา
    RegistrationNo: string; // หมายเลขทะเบียนยา

    StockQuantity: number; // จำนวนในสต็อก
    ReorderLevel: number; // ระดับเตือนสั่งซื้อใหม่
    PricePerUnit: number; // ราคาต่อหน่วย
    ImportDate: string; // วันที่นำเข้า
    ExpiryDate: string; // วันหมดอายุ

    Manufacturer: string; // ผู้ผลิต
    CountryOfOrigin: string; // ประเทศที่ผลิต
    BatchNumber: string; // ล็อตที่ผลิต

    SupplierID: number; // รหัสผู้จัดจำหน่าย
    Supplier: SupplierInterface; // เชื่อมต่อกับข้อมูลผู้จัดจำหน่าย

    UsageInstructions: string; // คำแนะนำการใช้ยา
    Indications: string; // ข้อบ่งใช้
    Contraindications: string; // ข้อห้ามใช้
    SideEffects: string; // ผลข้างเคียง

    DrugImage: string; // รูปภาพของยา
    Barcode: string; // รหัสบาร์โค้ด  
  }