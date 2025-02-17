import React, { useEffect, useState } from 'react';
import { Input, Form, Avatar, Button} from 'antd';
import './saveahistory.css'; // ไฟล์ CSS สำหรับการตกแต่ง
import { useParams } from 'react-router-dom';
import { TakeAHistoryInterface } from '../../../interface/ITakeAHistory';
import { PatientInterface } from '../../../interface/IPatient';
import { apiUrl, GetPatientByNationId, GetTakeAHistoryById} from '../../../service/https';



const SaveTakeAHistory: React.FC = () => {
  const [form] = Form.useForm();
   
    const [patient, setPatient] = useState<PatientInterface | null>(null);
    const [takeahistory, setTakeAHistory] = useState<TakeAHistoryInterface | null>(null);
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
  

  //  const onFinish = async (values: TakeAHistoryInterface) => {
  //      // console.log(employeeId);
  //      console.log(patient)
  //      console.log("Form values:", values);
  //      const dateString = values.last_menstruation_date
  //      const date = new Date(`${dateString}T00:00:00Z`)
  //      // สร้างข้อมูลสำหรับอัปเดต Patient
   
  //      values.last_menstruation_date = date
  //      values.weight = Number(values.weight)
  //      values.height = Number(values.height)
  //      values.systolic_blood_pressure = Number(values.systolic_blood_pressure)
  //      values.diastolic_blood_pressure = Number(values.diastolic_blood_pressure)
  //      values.pulse_rate = Number(values.pulse_rate)
  //      values.patient_id = patient?.ID
  //      values.drink_alcohol = Boolean(values.drink_alcohol)
  //      values.smoking = Boolean(values.smoking)
  //      values.employee_id = 1
   
  //    };

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
      <div 
        style={{ 
          color: '#5752A7', 
          fontSize: '24px', 
          fontWeight: 'bold', 
          marginBottom: '20px', 
          textAlign: 'center', 
          width: '100%' }}
      >
          ซักประวัติคนไข้
      </div>

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

      {/* ส่วนข้อมูลเพิ่มเติม */}
      <div 
        style={{ 
          display: 'flex', 
          justifyContent: 'space-between', 
          width: '70%' }}
      >
        <div 
          style={{ 
            background: 'white', 
            borderRadius: '10px', 
            padding: '15px', 
            boxShadow: '0px 2px 5px rgba(0, 0, 0, 0.2)', 
            width: '48%',
            marginRight: '30px' }}
        >
          <h3>ข้อมูลทั่วไป</h3>
            <Form style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap' }}>
              <Form style={{ display: 'flex', flexDirection: 'column', gap: '5px' }}>
                <Form.Item label="น้ำหนัก (กก.)" name="weight" style={{ width: '100%' }}>
                  <div>{takeahistory?.weight}</div>
                </Form.Item>
                <Form.Item label="ส่วนสูง (ซม.)" name="height" style={{ width: '100%' }}> 
                  <div>{takeahistory?.height}</div>
                </Form.Item>
                <Form.Item label="ค่าความดันขณะหัวใจบีบตัว (mmHg)" name="systolic_blood_pressure" style={{ width: '100%' }}> 
                  <div>{takeahistory?.systolic_blood_pressure}</div>
                </Form.Item>
                <Form.Item label="ค่าความดันขณะหัวใจคลายตัว (mmHg)" name="diastolic_blood_pressure" style={{ width: '100%' }}> 
                  <div>{takeahistory?.diastolic_blood_pressure}</div>
                </Form.Item>
                <Form.Item label="อัตราการเต้นของหัวใจ (bpm)" name="pulse_rate" style={{ width: '100%' }}> 
                  <div>{takeahistory?.pulse_rate}</div>
                </Form.Item>
              </Form>
                {(() => {
                    if (patient?.Gender?.gender_name === "Male" ) {
                      return <>
                      <Form.Item label="ดื่ม" name="drink_alcohol" >
                        <div>{takeahistory?.drink_alcohol}</div>
                      </Form.Item></>
                    }else{
                      return <>
                        <Form.Item
                          label="ประจำเดือนครั้งล่าสุด"
                          name="last_menstruation_date"
                          >
                          <div>{takeahistory?.last_menstruation_date}</div>
                        </Form.Item>
                        <Form.Item
                          label="ดื่ม"
                          name="drink_alcohol"
                        >
                          <div>{takeahistory?.drink_alcohol}</div>
                        </Form.Item>
                        </>
                    }
                  })()}
              <Form.Item label="สูบบุหรี่" name="smoking" style={{ width: '100%' }}> 
                <div>{takeahistory?.smoking}</div>
              </Form.Item>
              <Form.Item label="โรคประจำตัว" name="disease_name" style={{ width: '100%' }}> 
                <div>{patient? DateFormat:""}</div>
              </Form.Item>
            </Form>
          {/* เพิ่มข้อมูลทั่วไปที่นี่ */}
        </div>
        <div 
          style={{ 
            background: 'white', 
            borderRadius: '10px', 
            padding: '15px', 
            boxShadow: '0px 2px 5px rgba(0, 0, 0, 0.2)', 
            width: '48%' }}
        >
          <h3>อาการเบื้องต้น</h3>
            <Form style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap' }}>
              <Form.Item name="preliminary_symptoms" style={{ width: '100%' }}> 
                <div>{takeahistory?.preliminary_symptoms}</div>
              </Form.Item>
              
            </Form>
          {/* เพิ่มข้อมูลอาการเบื้องต้นที่นี่ */}
        </div>
      </div>

      {/* ปุ่ม */}
      <div 
        style={{  
          display: 'flex', 
          gap: '10px', 
          marginTop: '20px' }}
      >
        <Button type="primary" htmlType="submit" style={{ marginRight: '10px', border: '1px solid #5752A7' , backgroundColor: '#5752A7',padding: '10px 20px' , borderRadius: '5px', fontSize: '16px',}}>
          แก้ไขประวัติการรักษา
        </Button>
        <Button htmlType="button" style={{ marginRight: '10px', border: '1px solid #5752A7' ,backgroundColor: 'white',padding: '10px 20px' , borderRadius: '5px', fontSize: '16px', color: '#5752A7'}} onClick={() => form.resetFields()}>
          บันทึก
        </Button>
      </div>
    </div>
  );
};

export default SaveTakeAHistory;