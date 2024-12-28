export interface TakeAHistoryInterface {
  ID?: number; // ID อาจไม่จำเป็นต้องกำหนดในตอนสร้าง
  weight?: number; // น้ำหนัก
  height?: number; // ส่วนสูง
  preliminary_symptoms?: string; // อาการเบื้องต้น
  systolic_blood_pressure?: number; // ความดันโลหิตช่วงบน
  diastolic_blood_pressure?: number; // ความดันโลหิตช่วงล่าง
  pulse_rate?: number; // อัตราชีพจร
  smoking?: boolean; // การสูบบุหรี่
  drink_alcohol?: boolean; // การดื่มแอลกอฮอล์ 
  // DATE?: Date; // วันที่บันทึก
  // QueueStatus?: string;
  last_menstruation_date?: Date; // วันที่มีประจำเดือนครั้งสุดท้าย
  patient_id?: number;
  employee_id?: number;
  appointment_id?: number;
  disease_name?:  number;
  // MedicalRecordsID?: number | null; // ID ของ MedicalRecord
}


export interface Iupdatepatientdisease{
	id?:number;
	patient_id?:number;
	disease_id?:number[];
}