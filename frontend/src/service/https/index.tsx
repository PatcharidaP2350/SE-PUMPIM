import axios from "axios";
import { PatientInterface } from "../../interface/IPatient";
import { Iupdatepatientdisease, TakeAHistoryInterface } from "../../interface/ITakeAHistory";
import { AppointmentInterface } from "../../interface/IAppointment";

export const apiUrl = "http://localhost:8000"; // URL ของ Backend API
const Authorization = ("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImRvY3RvciIsImV4cCI6MTczNjYyNTQ3NSwiaXNzIjoiQXV0aFNlcnZpY2UifQ.i9aeHda_VvlXTwHRZRymkW8mY6N6HMwhjCcoKM29b6s");
// const Bearer = localStorage.getItem("token_type");  //ถ้ายังไม่เซต localStorage จะดึง Disease ไม่ได้


const requestOptions = {

  headers: {

    "Content-Type": "application/json",

    Authorization: `Bearer ${Authorization}`,

  },

};


// ฟังก์ชันสำหรับดึง token จาก localStorage
function getAuthHeaders() {
    return {
      "Content-Type": "application/json",
      Authorization: `Bearer ${Authorization}`,
    };
  }

// ฟังก์ชันสำหรับบันทึกข้อมูลผู้ป่วย
async function CreatePatient(data: PatientInterface) {

  return await axios

    .post(`${apiUrl}/patients`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


// ฟังก์ชันสำหรับดึงข้อมูลผู้ป่วยตาม ID
async function GetPatientById(id: number) {

  return await axios

    .get(`${apiUrl}/patients/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetPatientByNationId(nation_id: string) {

  return await axios

    .get(`${apiUrl}/patient/${nation_id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function ListPatients() {

  return await axios

    .get(`${apiUrl}/patients`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


async function DeletePatientByid(id: number) {

  return await axios

    .delete(`${apiUrl}/patients/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


// ฟังก์ชันสำหรับอัปเดตข้อมูลผู้ป่วยบางฟิลด์
async function UpdatePatientById(id: number, data: PatientInterface) {

  return await axios

    .put(`${apiUrl}/patients/${id}`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


// ฟังก์ชัน async สำหรับการสร้าง TakeAHistory
async function CreateTakeAHistory(data: TakeAHistoryInterface) {
  console.log("AAA",data)
  return await axios

    .post(`${apiUrl}/take_a_history`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function ListTakeAHistory() {

  return await axios

    .get(`${apiUrl}/take_a_history`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


// ฟังก์ชันสำหรับดึงข้อมูลผู้ป่วยตาม ID
async function GetTakeAHistoryById(id: number) {

  return await axios

    .get(`${apiUrl}/take_a_history/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


// ฟังก์ชันสำหรับอัปเดตข้อมูลผู้ป่วยบางฟิลด์
async function UpdateTakeAHistoryById(id: number, data: PatientInterface) {

  return await axios

    .put(`${apiUrl}/take_a_history/${id}`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


async function DeleteTakeAHistoryByid(id: number) {

  return await axios

    .delete(`${apiUrl}/take_a_history/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


async function GetPatientVisit() {

  return await axios

    .get(`${apiUrl}/patient_visit`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function getDurg() {
  const requestOptions = {
    //method: "GET",
    headers: getAuthHeaders(),
  };
  return await axios

  .get(`${apiUrl}/drugs`, requestOptions)

  .then((res) => res)

  .catch((e) => e.response);

}

async function getDiseases() {
  const requestOptions = {
    //method: "GET",
    headers: getAuthHeaders(),
  };
  return await axios

  .get(`${apiUrl}/diseases`, requestOptions)

  .then((res) => res)

  .catch((e) => e.response);
}


async function updatePatientDisease(data:Iupdatepatientdisease) {
  const requestOptions = {
    //method: "GET",
    headers: getAuthHeaders(),
  };
  return await axios

  .patch(`${apiUrl}/updatepatiendisease`, data, requestOptions)

  .then((res) => res)

  .catch((e) => e.response);
}

async function CreateAppointment(data: AppointmentInterface) {
  console.log("AAA",data)
  return await axios

    .post(`${apiUrl}/appointment`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetAppointmentById(id: number) {

  return await axios

    .get(`${apiUrl}/appointment/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


export {
    getAuthHeaders,
    CreatePatient,
    GetPatientById,
    GetPatientByNationId,
    ListPatients,
    DeletePatientByid,
    UpdatePatientById,
    CreateTakeAHistory,
    ListTakeAHistory,
    GetTakeAHistoryById,
    UpdateTakeAHistoryById,
    DeleteTakeAHistoryByid,
    GetPatientVisit,
    updatePatientDisease,
    CreateAppointment,
    GetAppointmentById,

    getDurg,
    getDiseases,

};
