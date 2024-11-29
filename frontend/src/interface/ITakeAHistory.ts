export interface ITakeAHistory {
    ID?: number; // ID อาจไม่จำเป็นต้องกำหนดในตอนสร้าง
    WEIGHT: number; // น้ำหนัก
    HEIGHT: number; // ส่วนสูง
    PRELIMINARY_SYMPTOMS: string; // อาการเบื้องต้น
    SYSTOLIC_BLOOD_PRESSURE: number; // ความดันโลหิตช่วงบน
    DIASTOLIC_BLOOD_PRESSURE: number; // ความดันโลหิตช่วงล่าง
    PULSE_RATE: number; // อัตราชีพจร
    SMOKING: string; // การสูบบุหรี่
    DRINK_ALCOHOL: string; // การดื่มแอลกอฮอล์
    DATE: Date; // วันที่บันทึก
    LAST_MENSTRUATION_DATE: Date; // วันที่มีประจำเดือนครั้งสุดท้าย
  
    // MEDICAL_RECORD_ID?: number | null; // ID ของ MedicalRecord
    // MEDICAL_RECORD?: IMedicalRecords | null; // ความสัมพันธ์กับ MedicalRecords
  }
  