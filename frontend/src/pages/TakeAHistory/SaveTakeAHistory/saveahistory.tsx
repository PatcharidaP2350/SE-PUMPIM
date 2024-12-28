import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom/client';
import { Layout, Menu, Dropdown, Input, Button, Form, Select, Row, Col, message, Avatar} from 'antd';
import { UserOutlined, VideoCameraOutlined, UploadOutlined } from '@ant-design/icons';
import './saveahistory.css'; // ไฟล์ CSS สำหรับการตกแต่ง
import avatarImage from '../../../../src/assets/pim.jpg'; // เปลี่ยนพาธให้ตรงกับไฟล์ของคุณ
import { useParams } from 'react-router-dom';
import { PatientVisitInterface } from '../../../interface/IPatientVisit';
import { TakeAHistoryInterface } from '../../../interface/ITakeAHistory';
import { PatientInterface } from '../../../interface/IPatient';
import { apiUrl, CreateTakeAHistory, getDiseases, getDurg, GetPatientById, GetPatientByNationId, GetTakeAHistoryById } from '../../../service/https';
import { DrugsInterface } from '../../../interface/IDrug';


const { Header, Sider, Content, Footer } = Layout;
const { Option } = Select;

const SaveTakeAHistory: React.FC = () => {
  const [search, setSearch] = useState<PatientInterface[]>([]);
  // const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  const [patient_visit, setPatientVisit] = useState<PatientVisitInterface[]>([]);
  const [take_a_history, setTakeAHistory] = useState<TakeAHistoryInterface | null>(null);
  const [form] = Form.useForm();
  const [diseases, setDiseases] = useState<{ ID: string; disease_name: string }[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [drawerForm] = Form.useForm(); // Create form instance for drawer
  const [history, setHistory] = useState<TakeAHistoryInterface | null>(null);
  const [patient, setPatient] = useState<PatientInterface | null>(null);
  const employeeId = Number(localStorage.getItem("id"));
  const { id } = useParams<{ id: any }>();
  const [drug, setdrug] = useState<{ ID: string; DrugName: string }[]>([]);

   const handleSearch = async(nation_id: string) => {
    const res = await GetPatientByNationId(nation_id);
    if (res){
      setPatient(res.data);
    } 
  };
  console.log(patient?.FirstName)

  const GetPatientId = async (patient_id: number) => {
    let res = await GetPatientById(patient_id);

    if (res.status === 200) {
      setPatient(res.data);
    } else {
      messageApi.open({
        type: "error",
        content: res.data.error,
      });
    }
  };

  const onFinish = async (values: TakeAHistoryInterface) => {
    console.log(values)
    let res = await CreateTakeAHistory(values, patient?.ID);  
    if (res.status == 201) {
      messageApi.open({
        type: "success",
        content: res.data.message,
      });
      setTimeout(function () {
      }, 2000);
    } else {
      messageApi.open({
        type: "error",
        content: res.data.error,
      });
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      if (!id) {
        message.error("Invalid ID");
        
        return;
      }

      try {
        setLoading(true);
        const result = await GetTakeAHistoryById(id);
        if (result.status === 200) {
          setHistory(result.data[0]);
          const drugsArray = result.data[0].DrugNames ? result.data[0].DrugNames.split(',') : [];
          drawerForm.setFieldsValue({
          drug: drugsArray // แบ่ง DrugNames ให้เป็น array
        });
        } else {
          message.error("ไม่สามารถดึงข้อมูลเวชระเบียนได้ สถานะ:" + result.status);
        }
      } catch (error) {
        console.error("Error fetching data:", error);
        message.error("เกิดข้อผิดพลาดในการดึงข้อมูล");
      } finally {
        setLoading(false);
      }
    };

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
      
      const fetchDrug = async () => {
        try {
          const response = await getDurg(); // Fetch diseases from the API
          if (response.status === 200) {
            setdrug(response.data);
          } else {
            message.error("ไม่สามารถดึงข้อมูลโรคได้ สถานะ: " + response.status);
          }
        } catch (error) {
          console.error("Error fetching diseases:", error);
          message.error("เกิดข้อผิดพลาดในการดึงข้อมูลโรค");
        }
      };
    fetchData();
    fetchDiseases();
    fetchDrug();
  }, [id]);

   // เมนูสำหรับ Dropdown
   const menu = (
    <Menu>
      <Menu.Item key="1">มีนัด</Menu.Item>
      <Menu.Item key="2">ไม่มีนัด</Menu.Item>
    </Menu>
  );

  return (
    <Layout 
    className="AddTakeAHistory"
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
            }}
          >
            ซักประวัติ
          </div>
          <div
            style={{
              width: "70%", // กำหนดความกว้างของกล่องค้นหา
              marginBottom: "-100px",
              marginLeft: "-450px",
             
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
              height: "30%", //ำหนดความสูงของกล่อง
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
                    src={patient?.PatientPicture!=null ? `${apiUrl}/${patient?.PatientPicture}` : "./Avatar/woman.png"} // URL รูปภาพที่คุณต้องการ
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
                    <div>{patient?.DateOfBirth}</div>
                  </Form.Item>
                </Col>
                <Col span={3} offset={1}>
                  <Form.Item label="กรุ๊ปเลือด" name="BloodGroupID" style={{ width: '100%' }}> 
                    <div>{patient?.BloodGroup?.blood_group}</div>
                  </Form.Item>
                </Col>
                <Col span={3} offset={1}>
                  <Form.Item label="อายุ" name="BloodGroupID" style={{ width: '100%' }}> 
                    <div>{patient?.BloodGroup?.blood_group}</div>
                  </Form.Item>
                </Col>
              </Row>
           </div>  
              <Form
                form={form}
                layout="vertical"
                onFinish={onFinish}
                style={{ padding: '5px' }}
              >
              {/* ส่วนกรอกข้อมูล */}
              <Row gutter={[10, 10]}>
                <Col span={3}>
                  <Form.Item label="น้ำหนัก (กก.)" name="Weight">
                    <div>{take_a_history?.Weight}</div>
                  </Form.Item>
                </Col>
                <Col span={3} offset={1}>
                  <Form.Item label="ส่วนสูง (ซม.)" name="Hight">
                    <div>{take_a_history?.Hight}</div>
                  </Form.Item>
                </Col>
              </Row>
              <Row gutter={[10, 10]}>
                <Col span={3}>
                  <Form.Item label="ค่าความดันขณะหัวใจบีบตัว (mmHg)" name="SystolicBloodPressure">
                    <div>{take_a_history?.SystolicBloodPressure}</div>
                  </Form.Item>
                </Col>
                <Col span={3} offset={1}>
                  <Form.Item label="ค่าความดันขณะหัวใจคลายตัว (mmHg)" name="DiastolicBloodPressure">
                    <div>{take_a_history?.DiastolicBloodPressure}</div>
                  </Form.Item>
                </Col>
                <Col span={3} offset={1}>
                  <Form.Item label="อัตราการเต้นของหัวใจ (bpm)" name="PulseRate">
                    <div>{take_a_history?.PulseRate}</div>
                  </Form.Item>
                </Col>
              </Row>
              <Row gutter={[10, 10]}>
                {/* ประจำเดือนครั้งล่าสุด */}
                <Col span={3}>
                  <Form.Item label="ประจำเดือนครั้งล่าสุด" name="LastMenstruationDate">
                    <div>{take_a_history?.PulseRate}</div>
                  </Form.Item>
                </Col>
                
                <Col span={3} offset={1}>
                  <Form.Item label="ดื่ม" name="DrinkAlcohol">
                    <div>{take_a_history?.DrinkAlcohol}</div>
                  </Form.Item>
                </Col>
                <Col span={3} offset={1}>
                  <Form.Item label="สูบบุหรี่" name="Smoking">
                    <div>{take_a_history?.Smoking}</div>
                  </Form.Item>
                </Col>
                {/* แพ้ยา */}
                <Col span={5} offset={1}>
                  <Form.Item label="ประวัติการแพ้ยา" name="DrugName">
                    <div>{take_a_history?.Smoking}</div>
                  </Form.Item>
                </Col>

                {/* โรคประจำตัว */}
                <Col span={5} offset={1}>
                  <Form.Item label="โรคประจำตัว" name="chronicDisease">
                    <div>{take_a_history?.Smoking}</div>
                  </Form.Item>
                </Col>
                <Col span={20}>
                  <Form.Item label="อาการเบื้องต้น" name="PreliminarySymptoms">
                    <div>{take_a_history?.PreliminarySymptoms}</div>
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
        </Content>
    </Layout>
  );
};

export default SaveTakeAHistory;