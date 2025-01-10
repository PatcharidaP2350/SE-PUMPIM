import { Col, DatePicker, TimePicker, Form } from "antd";


const CreateAppointment: React.FC = () => {
    
    return (
      <div
        className="CreateAppointment"
        style={{
          height: "100vh",
          margin: 0,
          padding: 0,
          width: "100%",
          
        }}
      >
        <div
            style={{
            fontSize: "24px", // ขนาดตัวอักษร
            fontWeight: "bold", // ทำให้ตัวหนา
            marginBottom: "20px", // ระยะห่างจากข้อความอื่นๆ
            textAlign: "center",
            }}
        >
            สร้างการนัด
        </div>
        <Col span={3} offset={1}>
            <Form.Item
                label="วันที่นัด"
                name="appointment_date"
                rules={[{ required: true, message: 'กรุณาเลือกวันที่นัด' }]}
            >
                <DatePicker 
                format="YYYY-MM-DD" 
                style={{ width: '100%' }} 
                placeholder="เลือกวันที่นัด" 
                />
            </Form.Item>
        </Col>
        <Col span={3}>
            <Form.Item
                label="เวลานัด"
                name="appointment_time"
                rules={[{ required: true, message: 'กรุณาเลือกเวลานัด' }]}
            >
                <TimePicker 
                format="HH:mm" 
                style={{ width: '100%' }} 
                placeholder="เลือกเวลานัด" 
                />
            </Form.Item>
        </Col>
      </div>
    );
  };
  
  export default CreateAppointment;