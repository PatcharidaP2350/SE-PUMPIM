export interface AppointmentInterface {
  id?: number; // ใช้สำหรับระบุ ID ของ Appointment (optional)
  appointment_date: string; // วันที่นัดหมาย (ในรูปแบบ ISO string)
  appointment_time: string; // เวลานัดหมาย (ในรูปแบบ ISO string)
  reason: string; // เหตุผลสำหรับการนัด
  status: string; // สถานะของการนัด
  note?: string; // หมายเหตุเพิ่มเติม (optional)
  employee_id: number; // รหัสพนักงานที่เกี่ยวข้องกับการนัด (required)
  medical_records_id?: number; // รหัสประวัติการรักษา (optional)
  // medical_records?: MedicalRecordsInterface; // ความสัมพันธ์กับ MedicalRecords (optional)
  // employee?: EmployeeInterface; // ความสัมพันธ์กับ Employee (optional)
}
