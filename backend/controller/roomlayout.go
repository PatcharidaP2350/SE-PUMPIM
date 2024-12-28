package controller

import (
    "net/http"
	"errors"
    "SE-B6527075/config"
	"SE-B6527075/entity"
    "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type (
    RoomLayoutResult struct{
        RoomLayoutID	uint `json:"room_layout_id"`
    
        // BuildingName  string `json:"building_name"`
    
        RoomNumber   string `json:"room_number"`
    
        RoomName string `json:"room_name"`
    
        PositionX int `json:"position_x"`
    
        PositionY int `json:"position_y"`
    
        // FloorNumber     string    `json:"floor_number"`
    };
    RoomInyput struct {
		FloorID  uint  `json:"floor_id"`
		BuildingID uint  `json:"building_id"`
		RoomNumber   string  `json:"room_number"`
		RoomName     string  `json:"room_name"`
		BedCapacity  int     `json:"bed_capacity"`
		PositionX    int     `json:"position_x"`
		PositionY    int     `json:"position_y"`
		PricePerNight float32 `json:"price_per_night"`
		DepartmentID uint `json:"department_id"`
		EmployeeID uint `json:"employee_id"`
	};
    RoomLayoutForEditResult struct{
        RoomLayoutID	uint `json:"room_layout_id"`
    
        RoomNumber   string `json:"room_number"`
    
        RoomName string `json:"room_name"`
    
        PositionX int `json:"position_x"`
    
        PositionY int `json:"position_y"`

        DepartmentID uint `json:"department_id"`

        EmployeeID uint `json:"employee_id"`

        BedCapacity  int     `json:"bed_capacity"`
    
    };


)


func GetRoomLayout(c *gin.Context) {
    var results []RoomLayoutResult
    buildingInput := c.Query("building_id")
    floorInput := c.Query("floor_id")

    db := config.DB()

    query := db.Model(&entity.RoomLayout{}).
        Select(`
            room_layouts.id AS room_layout_id,
            rooms.room_number,
            room_types.room_name,
            room_layouts.position_x,
            room_layouts.position_y
        `).
        Joins("INNER JOIN buildings ON floors.building_id = buildings.id").
        Joins("INNER JOIN floors ON room_layouts.floor_id = floors.id").
		Joins("INNER JOIN rooms ON room_layouts.room_id = rooms.id").
        Joins("INNER JOIN room_types ON rooms.room_type_id = room_types.id")
        

    // เพิ่มเงื่อนไขการค้นหาด้วย building_name และ floor_number (ถ้า floor_number มีค่า)
    if buildingInput != "" {
        query = query.Where("buildings.id LIKE ?", "%"+buildingInput+"%")
    }
    if floorInput != "" {
        query = query.Where("floors.id = ?", floorInput)
    }

    if err := query.Find(&results).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    if len(results) == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "ไม่พบข้อมูลผังห้องห้อง",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": results,
    })
}

func AddRoom(c *gin.Context) {
	var roomData RoomInyput;

	// ผูกข้อมูลจาก JSON ใน request มาใส่ใน roomData struct
	if err := c.ShouldBindJSON(&roomData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	// เริ่มเชื่อมต่อกับฐานข้อมูล
	db := config.DB()

	// ค้นหาอาคารและชั้นจากชื่อ
	var building entity.Building
	if err := db.Where("buildings.id = ?", roomData.BuildingID).First(&building).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบอาคาร",
		})
		return
	}

	var floor entity.Floor
	if err := db.Where("floors.id = ?", roomData.FloorID).First(&floor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบชั้น",
		})
		return
	}

    var roomType entity.RoomType
	if err := db.Where("room_name = ?", roomData.RoomName).First(&roomType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบประเภทห้อง",
		})
		return
	}

    // ใช้การทำธุรกรรมเพื่อเพิ่มข้อมูลทั้งหมด
	tx := db.Begin()

	// สร้างห้องใหม่
	room := entity.Room{
		RoomNumber:  roomData.RoomNumber,
		RoomTypeID:  roomType.ID, // ผูกกับ RoomType
		BedCapacity: roomData.BedCapacity,
		DepartmentID: roomData.DepartmentID,
		EmployeeID: roomData.EmployeeID,

	}

    if err := tx.Create(&room).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่สามารถสร้างห้องได้",
		})
		return
	}

// ตรวจสอบว่ามีข้อมูลในตำแหน่งที่ต้องการอยู่แล้วหรือไม่
var existingRoomLayout entity.RoomLayout
if err := tx.Joins("JOIN rooms ON rooms.id = room_layouts.room_id").
    Where(" room_layouts.floor_id = ? AND room_layouts.position_x = ? AND room_layouts.position_y = ?",
         floor.ID, roomData.PositionX, roomData.PositionY).
    First(&existingRoomLayout).Error; err == nil {
    // หากมีข้อมูลในตำแหน่งนี้อยู่แล้ว ให้แสดงข้อความหรือจัดการตามที่ต้องการ
    tx.Rollback()
    c.JSON(http.StatusBadRequest, gin.H{
        "error": "มีห้องในตำแหน่งนี้อยู่แล้ว กรุณาเลือกตำแหน่งใหม่",
    })
    return
} else if !errors.Is(err, gorm.ErrRecordNotFound) {
    // กรณีมีข้อผิดพลาดในการค้นหาอื่นๆ ให้ส่ง error กลับไป
    tx.Rollback()
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": "เกิดข้อผิดพลาดในการตรวจสอบตำแหน่ง",
    })
    return
}


	// สร้างข้อมูล RoomLayout
	roomLayout := entity.RoomLayout{
		FloorID: floor.ID,
		RoomID:     room.ID,
		PositionX:  roomData.PositionX,
		PositionY:  roomData.PositionY,
	}

	
	if err := tx.Create(&roomLayout).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่สามารถสร้างการจัดเรียงห้องได้",
		})
		return
	}

	// คอมมิตธุรกรรมหากทุกอย่างผ่านไปได้ด้วยดี
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "เกิดข้อผิดพลาดในการคอมมิตธุรกรรม",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ห้องถูกเพิ่มเรียบร้อยแล้ว",
        "data": room.ID,
	})
}

func DeleteRoom(c *gin.Context) {
    // รับ ID ของห้องที่ต้องการลบจาก URL parameter
    id := c.Param("id")
    
    // รับข้อมูล EmployeeID จาก JSON Request
    type DeleteInput struct {
        EmployeeID uint `json:"employee_id" binding:"required"`
    }
    var input DeleteInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "กรุณาระบุ EmployeeID",
        })
        return
    }

    // เชื่อมต่อฐานข้อมูล
    db := config.DB()

    // ตรวจสอบว่าห้องมีอยู่หรือไม่
    var room entity.Room
    if err := db.First(&room, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "ไม่พบห้องที่ต้องการลบ",
        })
        return
    }

    // ตรวจสอบว่าห้องมีอยู่หรือไม่
    var roomLayout entity.RoomLayout
    if err := db.First(&roomLayout, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "ไม่พบผังห้องที่ต้องการลบ",
        })
        return
    }

    // ใช้ Transaction เพื่อให้การลบทั้ง Room และ RoomLayout เป็น Atomic
    tx := db.Begin()

    // ลบ RoomLayout ที่เกี่ยวข้องกับ Room นี้
    if err := tx.Model(roomLayout).Updates(map[string]interface{}{
        "deleted_at": gorm.DeletedAt{Time: time.Now(), Valid: true},
    }).Error; err != nil {
    tx.Rollback()
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": "ไม่สามารถลบ RoomLayout ได้",
    })
    return
}


    // ทำ Soft Delete ห้องและบันทึก EmployeeID ของผู้ที่ลบ
    if err := tx.Model(&room).Updates(map[string]interface{}{
        "employee_id": input.EmployeeID,
        "deleted_at": gorm.DeletedAt{Time: time.Now(), Valid: true},
    }).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "ไม่สามารถลบห้องได้",
        })
        return
    }

    // Commit Transaction
    if err := tx.Commit().Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "ไม่สามารถ Commit การลบได้",
        })
        return
    }

    // ส่งผลลัพธ์กลับไปยัง Client
    c.JSON(http.StatusOK, gin.H{
        "message": "ลบห้องและ RoomLayout สำเร็จ",
    })
}

func GetDataForEditRoomLayout(c *gin.Context) {
    var results []RoomLayoutForEditResult
    buildingInput := c.Query("building_id")
    floorInput := c.Query("floor_id")

    db := config.DB()

    query := db.Model(&entity.RoomLayout{}).
        Select(`
            room_layouts.id AS room_layout_id,
            rooms.room_number,
            room_types.room_name,
            room_layouts.position_x,
            room_layouts.position_y,
            departments.id AS department_id,
            rooms.bed_capacity
        `).
        Joins("INNER JOIN floors ON room_layouts.floor_id = floors.id").
        Joins("INNER JOIN buildings ON floors.building_id = buildings.id").
        Joins("INNER JOIN rooms ON room_layouts.room_id = rooms.id").
        Joins("INNER JOIN room_types ON rooms.room_type_id = room_types.id").
        Joins("INNER JOIN departments ON rooms.department_id = departments.id")

    // เพิ่มเงื่อนไขการค้นหาด้วย building_name และ floor_number (ถ้า floor_number มีค่า)
    if buildingInput != "" {
        query = query.Where("buildings.id = ?", buildingInput)
    }
    if floorInput != "" {
        query = query.Where("floors.id = ?", floorInput)
    }

    if err := query.Find(&results).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    if len(results) == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "ไม่พบข้อมูลผังห้องตามเงื่อนไขที่กำหนด",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": results,
    })
}
