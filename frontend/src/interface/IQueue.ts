export interface IQueue {
    ID?: number; // ID ของคิว อาจไม่จำเป็นต้องกำหนดในตอนสร้าง
    QUEUE_NUMBER: string; // หมายเลขคิว
    QUEUE_DATE: string; // วันที่ของคิว
    QUEUE_TIME: Date; // เวลาของคิว
    STATUS: string; // สถานะของคิว
    PATIENT_ID: number; // ID ของผู้ป่วย
  }
  