// import { IGender } from "./IGender";

export interface IPatient {
    ID?: number;                // รหัสของผู้ป่วย
    NationID?: string;          // หมายเลขบัตรประชาชน
    FirstName?: string;         // ชื่อ
    LastName?: string;          // นามสกุล
    DateOfBirth?: string;       // วันเกิด
    Address?: string;           // ที่อยู่
    PhoneNumber?: string;       // เบอร์โทรศัพท์
    GenderID?: number;          // รหัสเพศ
    BloodGroupID?: number;      // รหัสกรุ๊ปเลือด
    AllergyInfo?: string;       // ข้อมูลการแพ้
    Diseases?: number[];        // ID ของโรคที่ผู้ป่วยมี
  }
  