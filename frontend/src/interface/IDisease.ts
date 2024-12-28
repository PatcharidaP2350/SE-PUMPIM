export interface IDisease {
    ID: number;
    disease_name: string;
    employees: any | null; // หรือระบุประเภทที่แน่นอนได้ถ้ามีโครงสร้างชัดเจน
    medicalrecords: any | null; // เช่นเดียวกัน
    patients: any | null; // เช่นเดียวกัน
  }
