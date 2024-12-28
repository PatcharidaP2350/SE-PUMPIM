package controller

import (
	"net/http"
	//"errors"
    "SE-B6527075/config"
	"SE-B6527075/entity"
    "github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

type (
	resultRoomType struct{
		ID uint `json:"id"`
		RoomName string `json:"room_name"`
	};
)

func GetTypeRoom(c *gin.Context){
	var room_type []resultRoomType

	results := config.DB().Model(&entity.RoomType{}).Find(&room_type)

	
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": room_type})
}