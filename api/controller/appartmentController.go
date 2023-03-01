package controller

import (
	model "booking/api/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReqInput struct {
	Admin_ID    string `json:"admin_id" binding:"required"`
	HousingType string `json:"housing_type" binding:"required"`
}

func CreateAppartment(ctx *gin.Context) {

	var (
		req    ReqInput
		aprtmt = model.Appartment{}
		er     error
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Fatal(err)
	}

	if model.CheckAdmin(req.Admin_ID) {
		aprtmt.HousingType, er = model.GetHouseType(req.HousingType)
		if er != nil {
			log.Fatal(er)
		}
		aprtmt.CreatedAt = time.Now()
		if err := model.DB.Create(&aprtmt).Error; err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Appartmant success created"})

	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Access Denied. You`re not admin!"})
	}
}
