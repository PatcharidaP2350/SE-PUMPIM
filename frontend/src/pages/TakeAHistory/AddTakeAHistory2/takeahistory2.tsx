import React, { useEffect, useState } from 'react';
import { Layout, Input, Button, Form, Select, Row, Col, message, Avatar } from 'antd';
import './takeahistory2.css'; // ไฟล์ CSS สำหรับการตกแต่ง
import { useNavigate, useParams } from 'react-router-dom';
import { TakeAHistoryInterface } from '../../../interface/ITakeAHistory';
import { PatientInterface } from '../../../interface/IPatient';
import { apiUrl, CreateTakeAHistory, getDiseases, GetPatientByNationId,updatePatientDisease} from '../../../service/https';

const { Header, Content} = Layout;
const { Option } = Select;

const AddTakeAHistory2: React.FC = () => {
  // const navigate = useNavigate();
  const [form] = Form.useForm();
  const [diseases, setDiseases] = useState<{ ID: string; disease_name: string }[]>([]);
  const [patient, setPatient] = useState<PatientInterface | null>(null);
  const employeeId = Number(localStorage.getItem("id"));
  const { id } = useParams<{ id: any }>();
  const months = ["มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", 
  "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"]
  const Datef = patient?.DateOfBirth?.slice(8,10)
  const monthFormat = months[Number(patient?.DateOfBirth?.slice(5,7)) - 1]
  const yearFormat = [Number(patient?.DateOfBirth?.slice(0,4)) + 543]
  const DateFormat = `${Datef} ${monthFormat} ${yearFormat}`
 
  
  const [formData, setFormData] = useState({ Weight: 0, Height: 0 });
  const getPatientIDfromSearch = (patient: PatientInterface) => {
    if (patient) {
      return patient.ID;  // ดึงค่า ID ของ Patient
    }
    return null;  // หากไม่มีข้อมูลให้คืนค่า null
  };


  const handleSearch = async (nation_id: string) => {
    const res = await GetPatientByNationId(nation_id);
    if (res) {
      setPatient(res.data);
      const patientID = getPatientIDfromSearch(res.data);  // ดึง ID จาก Patient
      console.log(patientID);  // แสดงผล ID ของ Patient
    }
  };

  const onFinish = async (values: TakeAHistoryInterface) => {
    // console.log(employeeId);
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

    try {
      
      const response = await CreateTakeAHistory(values);
      console.log(response);

      if (response.status === 201) {
        message.success("การบันทึกข้อมูล TakeAHistory สำเร็จ!");

        if (patient?.ID) {
          if (Array.isArray(values.disease_name)) {
            const diseaseUpdateResponse = await updatePatientDisease({
              patient_id: patient.ID,
              disease_id: values.disease_name.map((disease: any) => disease), // ดึงเฉพาะ ID ของโรค
            });

            if (diseaseUpdateResponse.status === 200) {
              message.success("บันทึกข้อมูลโรคสําเร็จ!");
            } else {
              message.error("บันทึกข้อมูลโรคไม่สําเร็จ สถานะ: " + diseaseUpdateResponse.status);
            }
          } else {
            console.error('values.disease_name is not an array');
          }
        }
        
      } else {
        throw new Error(`การบันทึกข้อมูล TakeAHistory ไม่สำเร็จ สถานะ: ${response.status}`);
      }

    } catch (error) {
      console.error("Error submitting medical record:", error);
      message.error("เกิดข้อผิดพลาดในการบันทึกข้อมูล");
    }
  };

  useEffect(() => {
    const fetchDiseases = async () => {
      try {
        const response = await getDiseases(); // Fetch diseases from the API
        if (response.status === 200) {
          setDiseases(response.data);
        } else {
          message.error("ไม่สามารถดึงข้อมูลโรคได้ สถานะ: " + response.status);
        }
      } catch (error) {
        console.error("Error fetching diseases:", error);
        message.error("เกิดข้อผิดพลาดในการดึงข้อมูลโรค");
      }
    };

    fetchDiseases();
  }, [id]);

  console.log(patient?.Gender?.gender_name)
  return (
    <Layout
      className="AddTakeAHistory2"
      style={{
        height: "100vh",
        backgroundColor: "#e2dfe4", // สีพื้นหลังของ Layout
        margin: 0,
        padding: 0,
        width: "100%",
      }}
    >
      <Header
        style={{
          backgroundColor: "#e2dfe4", // สีพื้นหลังของ Header
          color: "#5752A7",
          display: "flex", // ใช้ Flexbox
          alignItems: "center", // จัดตำแหน่งแนวตั้ง
          justifyContent: "space-between", // จัดตำแหน่งแบบซ้าย-ขวา
          fontSize: "24px",
          padding: "20px",
        }}
      >
        <div
          style={{
            marginBottom: "10px", // เพิ่มระยะห่างระหว่างข้อความและกล่องค้นหา
            marginLeft: "10px",
          }}
        >
          ซักประวัติ
        </div>
        <div
          style={{
            width: "70%", // กำหนดความกว้างของกล่องค้นหา
            marginBottom: "-100px",
            marginLeft: "-150px",
            marginRight: "300px",
          }}
        >
          <Input.Search
            placeholder="กรอกเลขประจำตัว"
            className="input-box"
            onSearch={(value) => handleSearch(value)} // ใส่ฟังก์ชันสำหรับค้นหา
          />
        </div>
        <div
          style={{
            marginRight: "20px",
            marginBottom: "-50px",
          }}
        >

        </div>
      </Header>
      <Content
        style={{
          display: "flex",
          justifyContent: "center", // จัดตำแหน่งกล่องให้อยู่กลางหน้าจอในแนวนอน
          alignItems: "flex-end", // จัดตำแหน่งกล่องให้อยู่กลางหน้าจอในแนวตั้ง
          backgroundColor: "#e2dfe4", // สีพื้นหลังของ Content
          height: "calc(100vh - 120px)", // คำนวณความสูงโดยลบ Header และ Footer
          padding: "20px",
        }}
      >
        <div
          style={{
            backgroundColor: "white", // สีพื้นหลังของกล่อง
            width: "99.5%", // กำหนดความกว้างของกล่อง
            height: "96%", // กำหนดความสูงของกล่อง
            padding: "30px", // ระยะห่างด้านใน
            boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)", // เพิ่มเงาให้ดูสวยงาม
          }}
        >
          {/* ข้อมูลพื้นฐาน */}
          <Row gutter={[10, 10]}>
            <Col
              span={2}
              style={{
                display: "flex",
                justifyContent: "flex-end",
              }}
            >
              <Avatar
                src={patient?.PatientPicture != null ? `${apiUrl}/${patient?.PatientPicture}` : (patient?.Gender?.gender_name === "Male" ? "./Avatar/man.png" : "./Avatar/woman.png")} // URL รูปภาพที่คุณต้องการ
                alt="avatar"
                size={70} // ขนาดของ Avatar
                style={{

                }}
              />
            </Col>
            <Col span={2}>
              <Form.Item
                label="ชื่อ"
                name="FirstName"
                style={{ width: '100%' }}>
                <div style={{ display: 'flex', alignItems: 'center' }}>
                  {patient?.FirstName}
                </div>
              </Form.Item>
            </Col>
            <Col span={3} offset={1}>
              <Form.Item label="นามสกุล" name="LastName" style={{ width: '100%' }}>
                <div>{patient?.LastName}</div>
              </Form.Item>
            </Col>
            <Col span={3} offset={1}>
              <Form.Item label="เพศ" name="GenderID" style={{ width: '100%' }}>
                <div>{patient?.Gender?.gender_name}</div>
              </Form.Item>
            </Col>
            <Col span={3} offset={1}>
              <Form.Item label="เกิดวันที่" name="DateOfBirth" style={{ width: '100%' }}>
                <div>{patient? DateFormat:""}</div>
              </Form.Item>
            </Col>
            <Col span={3} offset={1}>
              <Form.Item label="กรุ๊ปเลือด" name="BloodGroupID" style={{ width: '100%' }}>
                <div>{patient?.BloodGroup?.blood_group}</div>
              </Form.Item>
            </Col>
          </Row>
          <Form
            form={form}
            layout="vertical"
            onFinish={onFinish}
            style={{ padding: '5px' }}
          >
            {/* ส่วนกรอกข้อมูล */}
            <Row gutter={[10, 10]}>
              <Col span={3}>
                <Form.Item label="น้ำหนัก (กก.)" 
                  name="weight"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}>
                  <Input type="number" placeholder="กรอกน้ำหนัก"
                    onChange={(e) => setFormData({ ...formData, Weight: Number(e.target.value) })}  //------- เพิ่มเข้ามา -------- //
                  />
                </Form.Item>
              </Col>
              <Col span={3} offset={1}>
                <Form.Item
                  label="ส่วนสูง (ซม.)"
                  name="height"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}>
                  <Input type="number" placeholder="กรอกส่วนสูง"
                    onChange={(e) => setFormData({ ...formData, Height: Number(e.target.value) })}   //------- เพิ่มเข้ามา -------- //
                  />
                </Form.Item>
              </Col>
            </Row>
            <Row gutter={[10, 10]}>
              <Col span={3}>
                <Form.Item
                  label="ค่าความดันขณะหัวใจบีบตัว (mmHg)"
                  name="systolic_blood_pressure"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}>
                  <Input type="number" placeholder="กรอกค่าความดัน" />
                </Form.Item>
              </Col>
              <Col span={3} offset={1}>
                <Form.Item
                  label="ค่าความดันขณะหัวใจคลายตัว (mmHg)"
                  name="diastolic_blood_pressure"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}>
                  <Input type="number" placeholder="กรอกค่าความดัน" />
                </Form.Item>
              </Col>
              <Col span={3} offset={1}>
                <Form.Item
                  label="อัตราการเต้นของหัวใจ (bpm)"
                  name="pulse_rate"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}>
                  <Input type="number" placeholder="กรอกอัตราการเต้นของหัวใจ" />
                </Form.Item>
              </Col>
            </Row>
            <Row gutter={[10, 10]}>
              {(() => {
              if (patient?.Gender?.gender_name === "Male" ) {
                return <>
                <Col span={3}>
                <Form.Item
                  label="ดื่ม"
                  name="drink_alcohol"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}
                >
                  <Select placeholder="เลือกใช่หรือไม่">
                    <Select.Option value={1}>ใช่</Select.Option>
                    <Select.Option value={0}>ไม่ใช่</Select.Option>
                  </Select>
                </Form.Item>
                </Col></>
              }else{
                return <>
                <Col span={3}>
                  <Form.Item
                    label="ประจำเดือนครั้งล่าสุด"
                    name="last_menstruation_date"
                    rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}>
                    <Input type="date" />
                  </Form.Item>
                </Col>

                <Col span={3} offset={1}>
                <Form.Item
                  label="ดื่ม"
                  name="drink_alcohol"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}
                >
                  <Select placeholder="เลือกใช่หรือไม่">
                    <Select.Option value={1}>ใช่</Select.Option>
                    <Select.Option value={0}>ไม่ใช่</Select.Option>
                  </Select>
                </Form.Item>
                </Col>
                </>
              }

            })()}
              
              <Col span={3} offset={1}>
              <Form.Item
                label="สูบบุหรี่"
                name="smoking"
                rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}
              >
                <Select placeholder="เลือกใช่หรือไม่">
                  <Select.Option value={1}>ใช่</Select.Option>
                  <Select.Option value={0}>ไม่ใช่</Select.Option>
                </Select>
              </Form.Item>
              </Col>
              

              {/* โรคประจำตัว */}
              <Col span={5} offset={1}>
                <Form.Item
                  label="โรคประจำตัว"
                  name="disease_name"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}>
                  <Select
                    mode="tags" // ใช้ tags ให้ผู้ใช้พิมพ์หรือเลือกได้
                    allowClear
                    options={diseases.map((d) => ({
                      label: d.disease_name,
                      value: d.ID,
                    }))}
                    onChange={(value) => console.log("Updated disease allergies:", value)}
                  />
                </Form.Item>
              </Col>
              <Col span={20}>
                <Form.Item
                  label="อาการเบื้องต้น"
                  name="preliminary_symptoms"
                  rules={[{ required: true, message: "กรุณาเลือกข้อมูล" }]}>
                  <Input.TextArea placeholder="ระบุอาการเบื้องต้น" rows={3} />
                </Form.Item>
              </Col>
            </Row>
            <div style={{ textAlign: 'right', marginTop: '-10px', paddingRight: '20px' }}>
              <Button type="primary" htmlType="submit" style={{ marginRight: '10px' }}>
                บันทึก
              </Button>
              <Button htmlType="button" onClick={() => form.resetFields()}>
                ยกเลิก
              </Button>
            </div>
          </Form>

        </div>
      </Content>
    </Layout>
  );
};

export default AddTakeAHistory2;