import { IPatient } from "../../interface/IPatient";

const apiUrl = "http://localhost:8000"; // URL ของ Backend API


// ฟังก์ชันสำหรับดึง token จาก localStorage
function getAuthHeaders() {
    const token = localStorage.getItem("authToken");
    return {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    };
  }


// ฟังก์ชันสำหรับบันทึกข้อมูลผู้ป่วย
async function createPatient(patient: IPatient) {
    console.log("----------CreatePatient----------");
    const requestOptions = {
        method: "POST",
        headers: getAuthHeaders(),
        body: JSON.stringify(patient),
      };
  
    const response = await fetch(`${apiUrl}/patients`, requestOptions) // ใช้ apiUrl ที่กำหนดไว้ด้านบน
      .then(async (response) => {
        const data = await response.json(); // อ่านข้อมูล JSON ที่ Backend ตอบกลับมา
        return { status: response.status, data }; // รวม status และ data
      })
      .catch((error) => {
        console.error("Error creating patient:", error); // จัดการข้อผิดพลาด
        return { status: null, data: null }; // กรณีเกิดข้อผิดพลาด
      });
  
    return response; // คืนค่าผลลัพธ์ {status, data}
}


// ฟังก์ชันสำหรับดึงข้อมูลผู้ป่วยตาม ID
async function getPatientByid(id: string) {
  const requestOptions = {
    method: "GET",
    headers: getAuthHeaders(),
  };

 // ดึงข้อมูลผู้ป่วยจาก API
  const response = await fetch(`${apiUrl}/patients/${id}`, requestOptions);

  // ตรวจสอบสถานะของคำตอบจาก API
  if (response.status === 200) {
    const data: IPatient = await response.json(); // อ่านข้อมูล JSON ของผู้ป่วย
    return data; // คืนค่าข้อมูลผู้ป่วย
  } else {
    console.error(`Failed to fetch patient with ID: ${id}`); // ถ้าหากสถานะไม่ใช่ 200
    return false; // คืนค่า false หากไม่พบผู้ป่วย
  }
}





// ฟังก์ชันสำหรับอัปเดตข้อมูลผู้ป่วยบางฟิลด์
async function updatePatient(id: string, updates: Partial<IPatient>) {
  const requestOptions = {
    method: "PATCH",
    headers: getAuthHeaders(),
    body: JSON.stringify(updates),
  };
  const response = await fetch(`${apiUrl}/patients/${id}`, requestOptions)
      .then(async (response) => {
          const data = await response.json(); // อ่านข้อมูล JSON ที่ Backend ตอบกลับมา
          return { status: response.status, data }; // รวม status และ data
      })
      .catch((error) => {
          console.error(`Error updating patient with ID: ${id}`, error); // จัดการข้อผิดพลาด
          return { status: null, data: null }; // กรณีเกิดข้อผิดพลาด
      });

  return response; // คืนค่าผลลัพธ์ {status, data}
}

export {
    getAuthHeaders,
    createPatient,
    getPatientByid,
    updatePatient,
  
};
