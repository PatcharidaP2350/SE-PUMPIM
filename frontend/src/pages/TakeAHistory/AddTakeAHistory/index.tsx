import React, { useState } from 'react';
import ReactDOM from 'react-dom/client';
import { Layout, Menu, theme, Input, Button } from 'antd';
import { UserOutlined, VideoCameraOutlined, UploadOutlined } from '@ant-design/icons';
import './index.css'; // ไฟล์ CSS สำหรับการตกแต่ง

const { Header, Content, Footer, Sider } = Layout;

const AddTakeAHistory = () => {
    //const { token: { colorBgContainer, borderRadiusLG } } = theme.useToken();
  
    const [patientName, setPatientName] = useState('');
    const [patientAge, setPatientAge] = useState('');
    const [patientDisease, setPatientDisease] = useState('');
  
    const handleInputChange = (
      e: React.ChangeEvent<HTMLInputElement>,
      setter: React.Dispatch<React.SetStateAction<string>>
    ) => {
      setter(e.target.value);
    };
  
    return (
      <Layout className="layout-container">  {/* ใช้ class "layout-container" */}
        {/* Sidebar */}
        <Sider width={250} className="sider"> {/* ใช้ class "sider" */}
          <div className="demo-logo-vertical" />
        </Sider>
  
        <Layout>
          {/* Header */}
          <Header className="header">  {/* ใช้ class "header" */}
            ซักประวัติคนไข้
          </Header>
  
          {/* Content */}
          <Content className="content">  {/* ใช้ class "content" */}
            <div
              style={{
                minHeight: 360,
                padding: 24,
              }}
            >
              <h2>กรอกข้อมูลผู้ป่วย</h2>
  
              <div style={{ marginBottom: 16 }}>
                <Input
                  placeholder="ชื่อผู้ป่วย"
                  value={patientName}
                  onChange={(e) => handleInputChange(e, setPatientName)}
                />
              </div>
  
              <div style={{ marginBottom: 16 }}>
                <Input
                  placeholder="อายุผู้ป่วย"
                  value={patientAge}
                  onChange={(e) => handleInputChange(e, setPatientAge)}
                />
              </div>
  
              <div style={{ marginBottom: 16 }}>
                <Input
                  placeholder="โรคที่เป็น"
                  value={patientDisease}
                  onChange={(e) => handleInputChange(e, setPatientDisease)}
                />
              </div>
            </div>
          </Content>
  
          {/* Footer */}
          <Footer className="footer" style={{ display: 'flex', justifyContent: 'flex-end' }}>
            <Button
              type="primary"
              onClick={() => alert('บันทึกข้อมูลแล้ว')}
              style={{ marginRight: '10px' }}
            >
              บันทึกข้อมูล
            </Button>
            <Button type="default">ยกเลิก</Button>
          </Footer>
        </Layout>
      </Layout>
    );
  };
  
  export default AddTakeAHistory;