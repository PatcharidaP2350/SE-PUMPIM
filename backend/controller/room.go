package controller

import (
    "net/http"
    "SE-B6527075/config"
	"SE-B6527075/entity"
    "github.com/gin-gonic/gin"
)

type UpdateInput struct {
	RoomNumber   string `json:"room_number"`
	RoomName     string `json:"room_name"`
	BedCapacity  int    `json:"bed_capacity"`
	DepartmentID uint   `json:"department_id"`
	EmployeeID   uint   `json:"employee_id"`
}

func UpdateRoom(c *gin.Context) {
    id := c.Param("id") // รับ id ของห้อง

    var input UpdateInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ข้อมูลไม่ถูกต้อง",
        })
        return
    }

    db := config.DB()

    // ตรวจสอบว่าห้องมีอยู่หรือไม่
    var room entity.Room
    if err := db.First(&room, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "ไม่พบห้องที่ต้องการอัปเดต",
        })
        return
    }

    // ตรวจสอบ room_name เพื่อดึง room_type_id
    var roomType entity.RoomType
    if err := db.Where("room_name = ?", input.RoomName).First(&roomType).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ไม่พบประเภทห้อง",
        })
        return
    }

    // อัปเดตข้อมูลในฐานข้อมูล
    room.RoomNumber = input.RoomNumber
    room.RoomTypeID = roomType.ID
    room.BedCapacity = input.BedCapacity
    room.DepartmentID = input.DepartmentID
    room.EmployeeID = input.EmployeeID

    if err := db.Save(&room).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "ไม่สามารถอัปเดตข้อมูลได้",
        })
        return
    }

    // ส่งผลลัพธ์กลับไปยัง client
    c.JSON(http.StatusOK, gin.H{
        "message": "อัปเดตข้อมูลสำเร็จ",
    })
}
