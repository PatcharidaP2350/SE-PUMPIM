package controller

import (
	"net/http"
	"time"

	"SE-B6527075/config"
	"SE-B6527075/entity" // ให้ใช้ path ที่เหมาะสมกับโปรเจกต์ของคุณ
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateQueue(c *gin.Context) {
    var queue entity.Queue

    // ตรวจสอบการเชื่อมต่อฐานข้อมูล
    db := config.DB()
    if db == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
        return
    }

    // ผูกข้อมูลจาก request body กับ struct Queue
    if err := c.ShouldBindJSON(&queue); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ตรวจสอบ PatientID
    if queue.PatientID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid PatientID"})
        return
    }

    // ตรวจสอบว่ามี Patient ที่สอดคล้องกับ PatientID หรือไม่
    var patient entity.Patient
    if err := db.First(&patient, queue.PatientID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found"})
        return
    }

    // กำหนด QueueTime และ QueueDate
    queue.QueueTime = time.Now()
    queue.QueueDate = queue.QueueTime.Format("2006-01-02")

    // สร้างหมายเลขคิวอัตโนมัติ
    queue.QueueNumber = generateQueueNumber(db)

    // Debugging: แสดงข้อมูล Queue ก่อนบันทึก
    fmt.Printf("Queue Data: %+v\n", queue)

    // บันทึกข้อมูลลงในฐานข้อมูล
    if err := db.Create(&queue).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Debugging: แสดงข้อความเมื่อสร้างสำเร็จ
    fmt.Println("Queue created successfully")

    // ส่ง response กลับ
    c.JSON(http.StatusOK, gin.H{"message": "สร้างคิวสำเร็จ", "queue": queue})
}

// ฟังก์ชันสำหรับสร้างหมายเลขคิวอัตโนมัติ
func generateQueueNumber(db *gorm.DB) string {
    var lastQueue entity.Queue
    db.Last(&lastQueue)

    // กำหนดหมายเลขคิวถัดไป
    lastQueueNumber := lastQueue.QueueNumber
    newQueueNumber := "Q-" + fmt.Sprintf("%05d", getNextQueueNumber(lastQueueNumber))

    return newQueueNumber
}

// ฟังก์ชันเพื่อดึงหมายเลขคิวถัดไปจากหมายเลขที่มีอยู่
func getNextQueueNumber(lastQueueNumber string) int {
    var lastNum int
    if len(lastQueueNumber) > 2 {
        fmt.Sscanf(lastQueueNumber[2:], "%d", &lastNum)
    }
    return lastNum + 1
}




// GetQueue - ฟังก์ชันสำหรับดึงข้อมูล Queue โดยใช้ QueueID
func GetQueue(c *gin.Context) {
	// ดึง QueueID จาก URL parameter
	queueID := c.Param("queue_id")
	db := config.DB()

	// ค้นหา Queue จากฐานข้อมูล
	var queue entity.Queue
	if err := db.First(&queue, queueID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Queue not found"})
		return
	}

	// ตอบกลับข้อมูล Queue ที่พบ
	c.JSON(http.StatusOK, gin.H{"data": queue})
}




// ListQueue - ฟังก์ชันสำหรับดึงข้อมูล Queue ทั้งหมด
func ListQueue(c *gin.Context) {
	db := config.DB()

	// ค้นหาทุก Queue จากฐานข้อมูล
	var queues []entity.Queue
	if err := db.Find(&queues).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve queues"})
		return
	}

	// ตอบกลับข้อมูล Queue ทั้งหมด
	c.JSON(http.StatusOK, gin.H{"data": queues})
}





//ไม่ Gen คิวให้

// func CreateQueue(c *gin.Context) {
//     var queue entity.Queue

//     // ตรวจสอบการเชื่อมต่อฐานข้อมูล
//     db := config.DB()
//     if db == nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
//         return
//     }

//     // ผูกข้อมูลจาก request body กับ struct Queue
//     if err := c.ShouldBindJSON(&queue); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // ตรวจสอบ PatientID
//     if queue.PatientID == 0 {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid PatientID"})
//         return
//     }

//     // ตรวจสอบว่ามี Patient ที่สอดคล้องกับ PatientID หรือไม่
//     var patient entity.Patient
//     if err := db.First(&patient, queue.PatientID).Error; err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found"})
//         return
//     }

//     // กำหนด QueueTime และ QueueDate
//     queue.QueueTime = time.Now()
//     queue.QueueDate = queue.QueueTime.Format("2006-01-02")

//     // Debugging: แสดงข้อมูล Queue ก่อนบันทึก
//     fmt.Printf("Queue Data: %+v\n", queue)

//     // บันทึกข้อมูลลงในฐานข้อมูล
//     if err := db.Create(&queue).Error; err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // Debugging: แสดงข้อความเมื่อสร้างสำเร็จ
//     fmt.Println("Queue created successfully")

//     // ส่ง response กลับ
//     c.JSON(http.StatusOK, gin.H{"message": "สร้างคิวสำเร็จ", "queue": queue})
// }


