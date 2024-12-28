package controller

import (
    "net/http"
    "SE-B6527075/config"
	"SE-B6527075/entity"
    "github.com/gin-gonic/gin"
)
type (
	Floor struct{
        ID uint `json:"id"`
		FloorNumber string `json:"floor_number"`
	};
    inputFloor struct{
		FloorNumber string `json:"floor_number"`
        BuildingID uint `json:"building_id"`
	};
    
)
	

func GetFloor(c *gin.Context) {
    building_ID := c.Param("building_id")
    var results []Floor

    db := config.DB()
    query := db.Model(&entity.Floor{}).
        Select(`floors.id As id,
                floors.floor_number AS floor_number`).
        Joins("INNER JOIN buildings ON buildings.id = floors.building_id")

    if building_ID != "" {
        query = query.Where("floors.building_id = ?", building_ID)
    }

    if err := query.Find(&results).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

     // ตรวจสอบว่ามีข้อมูลหรือไม่
     if len(results) == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "อาคารนี้ยังไม่มีชั้นในอาคาร",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": results} )
}

func AddFloor(c *gin.Context) {
    var input inputFloor

    // ผูกข้อมูล JSON เข้ากับโครงสร้าง Floor
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB()

    var building entity.Building
	if err := db.Where("id = ?", input.BuildingID).First(&building).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบอาคาร",
		})
		return
	}

     // ตรวจสอบว่ามี FloorNumber ซ้ำอยู่แล้วหรือไม่
     var existingFloor entity.Floor
     if err := db.Where("floor_number = ? ", input.FloorNumber).
         First(&existingFloor).Error; err == nil {
         c.JSON(http.StatusConflict, gin.H{
             "error": "ชั้นนี้มีอยู่ในระบบแล้ว",
             "details": gin.H{
                 "floor_number": existingFloor.FloorNumber,
                 "building_id":  existingFloor.BuildingID,
             },
         })
         return
     }

    // ใช้การทำธุรกรรมเพื่อเพิ่มข้อมูลทั้งหมด
	tx := db.Begin()

    floors := entity.Floor{
        FloorNumber: input.FloorNumber,
        BuildingID:  input.BuildingID,
    }

    if err := tx.Create(&floors).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่สามารถสร้างชั้นได้ได้",
		})
		return
	}

    // คอมมิตการทำธุรกรรม
	tx.Commit()

    c.JSON(http.StatusOK, gin.H{"message": "เพิ่มสำเร็จ", "floor": floors.ID})
}
