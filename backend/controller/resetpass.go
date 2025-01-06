package controller

import (
	"fmt"
	"net/http"
	"time"
	"math/rand"

	"SE-B6527075/config"
	"SE-B6527075/entity"
	"SE-B6527075/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	
)

func ResetPasswordController(c *gin.Context) {
	type RequestPayload struct {
		Email string `json:"email" binding:"required"`
	}

	var payload RequestPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณากรอกอีเมลให้ถูกต้อง"})
		return
	}

	db := config.DB()

	// ตรวจสอบว่า Email มีในระบบหรือไม่
	var employee entity.Employee
	if err := db.Where("email = ?", payload.Email).First(&employee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบอีเมลในระบบ"})
		return
	}

	// สร้าง UUID และตั้งค่าเวลาหมดอายุ 5 นาที
	resetToken, err := Generate6DigitToken(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้างโทเค็นรีเซ็ตรหัสผ่านได้"})
		return
	}
	//resetExpiry := time.Now().Add(15 * time.Minute)
	// resetExpiry := time.Now().Add(5 * time.Minute)
	// บันทึก ResetToken และ ResetTokenExpiry ลงในฐานข้อมูล
	// employee.ResetToken = resetToken
	// employee.ResetTokenExpiry = resetExpiry
	if err := db.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถบันทึกโทเค็นรีเซ็ตรหัสผ่านได้"})
		return
	}

	// ส่งอีเมล
	// ส่งอีเมล
	subject := "โทเค็นสำหรับการรีเซ็ตรหัสผ่านของคุณ"
body := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="th">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Reset Password</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f6f9fc;
            color: #333;
            line-height: 1.8;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 600px;
            margin: 40px auto;
            background: #ffffff;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            border: 1px solid #e0e0e0;
        }
        .header {
            text-align: center;
            margin-bottom: 30px;
        }
        .header h1 {
            font-size: 24px;
            color: #1a73e8;
            margin: 0;
        }
        .content p {
            margin: 0 0 15px;
            font-size: 16px;
            color: #4a4a4a;
        }
        .token-boxes {
            display: flex;
            justify-content: center;
            margin: 20px 0;
        }
        .token-card {
            text-align: center;
            font-size: 22px;
            color: #ffffff;
            background-color: #1a73e8;
            padding: 15px 20px;
            border-radius: 6px;
            margin: 0 5px;
            font-weight: bold;
            width: 40px;
            height: 40px;
        }
        .footer {
            font-size: 14px;
            color: #777;
            text-align: center;
            margin-top: 30px;
            border-top: 1px solid #e0e0e0;
            padding-top: 20px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>รีเซ็ตรหัสผ่านของคุณ</h1>
        </div>
        <div class="content">
            <p>สวัสดีค่ะ/ครับ,</p>
            <p>นี่คือโทเค็นสำหรับการรีเซ็ตรหัสผ่านของคุณ:</p>
            <div class="token-boxes">
                %s
            </div>
            <p>กรุณาใช้โทเค็นนี้เพื่อรีเซ็ตรหัสผ่านของคุณภายใน 15 วินาที</p>
            <p>หากคุณไม่ได้ร้องขอการรีเซ็ตรหัสผ่าน กรุณาเพิกเฉยต่ออีเมลนี้</p>
            <p>ขอบคุณค่ะ/ครับ,<br>ทีมงานของเรา</p>
        </div>
        <div class="footer">
            <p>คุณได้รับอีเมลฉบับนี้เนื่องจากคุณเป็นสมาชิกในระบบของเรา</p>
            <p>หากมีคำถามเพิ่มเติม กรุณาติดต่อเราได้ตลอดเวลา</p>
        </div>
    </div>
</body>
</html>
`, formatTokenIntoCards(resetToken))



	// เรียกใช้ฟังก์ชันส่งอีเมล
	if err := SendEmail(payload.Email, subject, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถส่งอีเมลได้"})
		return
	}

	// ส่ง Response
	c.JSON(http.StatusOK, gin.H{
		"message": "ระบบได้ส่งโทเค็นไปยังอีเมลของคุณแล้ว",
	})
}


// ฟังก์ชันที่จะแปลงโทเค็นให้เป็นการ์ดสำหรับแต่ละตัวเลข
func formatTokenIntoCards(token string) string {
	var cardsHTML string

	// loop ผ่านตัวอักษรในโทเค็น และสร้างการ์ดสำหรับแต่ละตัวเลข
	for _, char := range token {
		// สร้างกล่องการ์ดสำหรับตัวเลข
		cardsHTML += fmt.Sprintf(`<div class="token-card">%c</div>`, char)
	}

	return cardsHTML
}


func Generate6DigitToken(db *gorm.DB) (string, error) {
	for {
		// สุ่มเลข 6 หลัก
		rand.Seed(time.Now().UnixNano())
		token := fmt.Sprintf("%06d", rand.Intn(1000000))

		// ตรวจสอบว่า token นี้มีอยู่ในฐานข้อมูลหรือไม่
		var existing entity.Employee
		if err := db.Where("reset_token = ?", token).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// ถ้าไม่เจอ record ที่มี token ซ้ำ ให้ return token
				return token, nil
			}
			// หากเกิดข้อผิดพลาดอื่นใน query ให้คืน error
			return "", err
		}
		// ถ้าเจอ token ซ้ำ ลูปสุ่มใหม่
	}
}


// ValidateResetTokenController ตรวจสอบว่า UUID ถูกต้องและยังไม่หมดอายุ
func ValidateResetTokenController(c *gin.Context) {
	type RequestPayload struct {
		Token string `json:"token" binding:"required"`
	}

	var payload RequestPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "คำขอไม่ถูกต้อง"})
		return
	}

	db := config.DB()

	// ตรวจสอบว่า Token มีในระบบและยังไม่หมดอายุ
	var employee entity.Employee
	if err := db.Where("reset_token = ?", payload.Token).First(&employee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "โทเค็นไม่ถูกต้องหรือหมดอายุ"})
		return
	}

	// ตรวจสอบเวลาหมดอายุ
	// if time.Now().After(employee.ResetTokenExpiry) {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "โทเค็นหมดอายุแล้ว"})
	// 	return
	// }

	// สร้าง JWT Token
	jwtWrapper := services.JwtWrapper{
		SecretKey:       config.GetSecretKey(), // ใช้คีย์จาก config
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	// ใช้ Username ของ Employee เพื่อสร้าง Token
	tokenString, err := jwtWrapper.GenerateToken(employee.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้างโทเค็นได้"})
		return
	}

	// ส่ง Response กลับ พร้อม JWT Token และ Employee ID
	c.JSON(http.StatusOK, gin.H{
		"message": "โทเค็นถูกต้อง",
		"jwt":     tokenString, // JWT Token ที่ส่งกลับ
		"id":      employee.ID, // ส่ง ID ของ Employee กลับไป
	})
}

func CleanExpiredTokens() {
	db := config.DB()

	// ล้าง Token ที่หมดอายุ
	if err := db.Model(&entity.Employee{}). // เพิ่ม "{}" หลัง "Employee" เพื่อบอกว่าเป็นอินสแตนซ์เปล่า
						Where("reset_token_expiry < ?", time.Now()).
						Updates(map[string]interface{}{
			"reset_token":        "",
			"reset_token_expiry": nil,
		}).Error; err != nil {
		//fmt.Println("เกิดข้อผิดพลาดในการล้างโทเค็นที่หมดอายุ:", err)
	} else {
		//fmt.Println("ล้างโทเค็นที่หมดอายุสำเร็จ")
	}
}

// StartCleanupJob ตั้ง Cron Job เพื่อล้าง UUID ที่หมดอายุ

