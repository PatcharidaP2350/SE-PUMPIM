package main

import (
	"net/http"

	"SE-B6527075/config"
	"SE-B6527075/controller"
	"SE-B6527075/middlewares"

	"github.com/gin-gonic/gin"
)

const PORT = "8000"

func main() {

	// เปิดการเชื่อมต่อฐานข้อมูล
	config.ConnectionDB()

	// สร้างตารางฐานข้อมูล
	config.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// // เส้นทางที่ไม่ต้องตรวจสอบโทเค็น เช่นการลงชื่อเข้าใช้
	// r.POST("/auth/signin", controller.EmployeeSignin)

	// r.POST("/reset-password", controller.ResetPasswordController)
	// r.POST("/validate-reset-token", controller.ValidateResetTokenController)

	// กลุ่มเส้นทางที่ต้องตรวจสอบโทเค็น
	protected := r.Group("/")
	protected.Use(middlewares.Authorizes()) // เรียกใช้ Authorizes middleware เพื่อเช็คโทเค็น
	{
		// เส้นทาง CreatePatient
		protected.POST("/patients", controller.CreatePatient)
		// เส้นทาง GetPatient
		protected.GET("/patients/:id", controller.GetPatient) // เพิ่มเส้นทาง GetPatient
		// เส้นทาง ListPatient
		protected.GET("/patients", controller.ListPatient) // เพิ่มเส้นทาง ListPatient
		// เส้นทาง DeletePatient
		protected.DELETE("/patients/:id", controller.DeletePatient) // เพิ่มเส้นทาง DeletePatient
		// เส้นทาง UpdatePatient
		protected.PATCH("/patients/:id", controller.UpdatePatient) // เพิ่มเส้นทาง UpdatePatient

		// เส้นทาง CreateQueue
		protected.POST("/queues", controller.CreateQueue)
		// เส้นทาง GetQueue
		protected.GET("/queues/:id", controller.GetQueue)
		// เส้นทาง ListQueue
		protected.GET("/queues", controller.ListQueue)

		// เส้นทาง CreateTakeAHistory
		protected.POST("/take_a_history", controller.CreateTakeAHistory)
		// เส้นทาง ListTakeAHistories
		protected.GET("/take_a_history", controller.ListTakeAHistory)
		// เส้นทาง GetTakeAHistory
		protected.GET("/take_a_history/:id", controller.GetTakeAHistory)
		// เส้นทาง UpdateTakeAHistory
		protected.PATCH("/take_a_history/:id", controller.UpdateTakeAHistory)
		// เส้นทาง DeleteTakeAHistory
		protected.DELETE("/take_a_history/:id", controller.DeleteTakeAHistory) // เพิ่มเส้นทาง DeleteTakeAHistory


	}

	// เส้นทางสำหรับตรวจสอบสถานะของ API
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	// เริ่มต้นเซิร์ฟเวอร์
	r.Run("localhost:" + PORT)
}

// ฟังก์ชัน CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
