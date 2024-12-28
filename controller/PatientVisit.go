package controller

import (
	"net/http"


	"SE-B6527075/entity"
	"SE-B6527075/config"
	"github.com/gin-gonic/gin"
)

// GET 
func GetPatientVisit(c *gin.Context) {
	var patient_visit []entity.PatientVisit

	db := config.DB()

	db.Find(&patient_visit)
	c.JSON(http.StatusOK, &patient_visit)
}
