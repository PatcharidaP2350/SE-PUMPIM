import React, { useEffect, useState } from 'react';
import { Input, Form, Avatar} from 'antd';
import './createpatient.css'; // ไฟล์ CSS สำหรับการตกแต่ง
import { useParams } from 'react-router-dom';
import { TakeAHistoryInterface } from '../../../interface/ITakeAHistory';
import { PatientInterface } from '../../../interface/IPatient';
import { apiUrl,   CreatePatient,   GetPatientByNationId} from '../../../service/https';



const AddPatient: React.FC = () => {
  const [form] = Form.useForm();
   
    const [patient, setPatient] = useState<PatientInterface | null>(null);
    // const employeeId = Number(localStorage.getItem("id"));
    const { id } = useParams<{ id: any }>();
    const months = ["มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", 
    "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"]
    const Datef = patient?.DateOfBirth?.slice(8,10)
    const monthFormat = months[Number(patient?.DateOfBirth?.slice(5,7)) - 1]
    const yearFormat = [Number(patient?.DateOfBirth?.slice(0,4)) + 543]
    const DateFormat = `${Datef} ${monthFormat} ${yearFormat}`
    
    const [formData, setFormData] = useState({ Weight: 0, Height: 0 });
      const getPatientIDfromSearch = (patient: PatientInterface) => {
        console.log(patient);
        if (patient) {
          return patient.ID;  // ดึงค่า ID ของ Patient
        }
        return null;  // หากไม่มีข้อมูลให้คืนค่า null
      };
    
   const handleSearch = async (nation_id: string) => {
    console.log(nation_id)
       const res = await GetPatientByNationId(nation_id);
       if (res) {
         setPatient(res.data);
         console.log(res.data.id);
         const patientID = getPatientIDfromSearch(res.data);  // ดึง ID จาก Patient
         console.log(patientID);  // แสดงผล ID ของ Patient
         console.log(patient?.ID);  // แสดงผล ID ของ Patient
       }
     };
  
   const onFinish = async (values: TakeAHistoryInterface) => {
       // console.log(employeeId);
       console.log(patient)
       console.log("Form values:", values);
       const dateString = values.last_menstruation_date
       const date = new Date(`${dateString}T00:00:00Z`)
       // สร้างข้อมูลสำหรับอัปเดต Patient
   
       values.last_menstruation_date = date
       values.weight = Number(values.weight)
       values.height = Number(values.height)
       values.systolic_blood_pressure = Number(values.systolic_blood_pressure)
       values.diastolic_blood_pressure = Number(values.diastolic_blood_pressure)
       values.pulse_rate = Number(values.pulse_rate)
       values.patient_id = patient?.ID
       values.drink_alcohol = Boolean(values.drink_alcohol)
       values.smoking = Boolean(values.smoking)
       values.employee_id = 1
   
     };

  useEffect(() => {
    }, [id]);

  return (
    <div 
      style={{ 
        backgroundColor: '#e2dfe4', 
        padding: '20px', display: 'flex', 
        justifyContent: 'center', 
        flexDirection: 'column', 
        alignItems: 'center', 
        fontFamily: 'Arial, sans-serif' }}
    >

      {/* กล่องค้นหา */}
      <div 
        style={{ 
          width: '70%', 
          marginBottom: '20px' }}
      >
        <Input.Search
          placeholder="กรอกเลขประจำตัวประชาชน"
          onSearch={handleSearch}
        />
      </div>

      {/* ข้อมูลคนไข้ */}
      <div 
        style={{ 
          display: 'flex', 
          justifyContent: 'center', 
          width: '100%' }}
      >
        <div 
          style={{ 
            display: 'flex', 
            alignItems: 'center', 
            background: 'white', 
            borderRadius: '10px', 
            padding: '15px', 
            boxShadow: '0px 2px 5px rgba(0, 0, 0, 0.2)', 
            width: '68%', 
            marginBottom: '20px' }}
        >
          <Avatar
            src={patient?.PatientPicture != null ? `${apiUrl}/${patient?.PatientPicture}` : (patient?.Gender?.gender_name === "Male" ? "./Avatar/man.png" : "./Avatar/woman.png")}  // URL รูปภาพที่คุณต้องการ
            alt="avatar"
            size={70} // ขนาดของ Avatar
            style={{
            marginLeft: '50px',
            marginTop: '-10px'
            }}
          />
          <div 
            style={{ paddingLeft: '0px' }}
          >
            <Form style={{ display: 'flex', flexDirection: 'row', gap: '10px' }}>

              <Form.Item label="ชื่อ" name="FirstName" style={{ width: '100%' }}><div style={{ display: 'flex', alignItems: 'center' }}>{patient?.FirstName}</div>
              </Form.Item>
              <Form.Item label="นามสกุล" name="LastName" style={{ width: '100%' }}><div style={{ display: 'flex', alignItems: 'center' }}>{patient?.LastName}</div>
              </Form.Item>
              <Form.Item label="กรุ๊ปเลือด" name="BloodGroupID" style={{ width: '100%' }}> 
                <div>{patient?.BloodGroup?.blood_group}</div>
              </Form.Item>
            </Form>
            <Form style={{ display: 'flex', flexDirection: 'row', gap: '20px' }}>
              <Form.Item label="เพศ" name="GenderID" style={{ width: '100%' }}>
                <div>{patient?.Gender?.gender_name}</div>
              </Form.Item>
              <Form.Item label="เกิดวันที่" name="DateOfBirth" style={{ width: '100%' }}> 
                <div>{patient? DateFormat:""}</div>
              </Form.Item>
              <Form.Item label="อายุ" name="DateOfBirth" style={{ width: '100%' }}> 
                <div>{patient? DateFormat:""}</div>
              </Form.Item>
            </Form>
          </div>
        </div>
      </div>
      <div 
        style={{ 
          display: 'flex', 
          gap: '10px', 
          marginTop: '20px' }}
      >
        <button style={{ padding: '10px 20px', border: 'none', borderRadius: '5px', fontSize: '16px', cursor: 'pointer', backgroundColor: '#5752A7', color: 'white' }}>บันทึก</button>
        <button style={{ padding: '10px 20px', border: '1px solid #5752A7', borderRadius: '5px', fontSize: '16px', cursor: 'pointer', backgroundColor: 'white', color: '#5752A7' }}>แก้ไขประวัติการรักษา</button>
      </div>
    </div>
  );
};

export default AddPatient;