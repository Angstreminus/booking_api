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
	RentPrice   int    `json:"rent_price" binding:"required"`
	HousingType string `json:"housing_type" binding:"required"`
}

//TODO Time conversion - error!
//TODO Pars time in struct

type ReqBookBody struct {
	User_id        uint
	Appartment_id  uint
	Year_start     int
	Month_start    int
	Day_strart     int
	Year_end       int
	Month_end      int
	Day_end        int
	Internet       bool
	Food           bool
	LuggageStorage bool
	Delivery       bool
	ClothCleaner   bool
}

func BookHousing(ctx *gin.Context) {
	var (
		req    ReqBookBody
		appart = model.Appartment{}
		srv    = model.Service{}
		user   = model.User{}
		err    error
	)

	//!? 1. Find Appartment + set UserRefer = user_id
	//!? 2. Find User append this appartment to User Appartments slice
	//!? 3. Add date fields: Book date start - Book date end
	//!? 4. Create services and fill them

	if err = ctx.ShouldBindJSON(&req); err != nil {
		log.Fatal(err)
	}

	//!Check if the user exist + his role is Admin or User
	//! else deny request
	user, err = FindUsrById(req.User_id)
	if err != nil {
		log.Fatal(err)
	}
	//!Check if user has registrated
	if (user.Role != "Admin") && (user.Role != "User") {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Access Denied. You`re not allowed to book. Register first please!"})
	}

	appart, err = model.FindAppartmentById(req.Appartment_id)
	if err != nil {
		log.Fatal(err)
	}

	beginDate := time.Date(req.Year_start, time.Month(req.Month_start), req.Day_strart, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(req.Year_end, time.Month(req.Month_end), req.Day_end, 0, 0, 0, 0, time.UTC)

	//!check if already booked
	//!time.IsZero - Checks if it is Default Value
	if !appart.BookEndDate.IsZero() || beginDate.Before(appart.BookStartDate) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Access Denied. Already booked!"})
	}

	//!rewrinting appartment`s foreign key
	appart.UserRefer = user.ID

	//!creating appartment`s service pack
	srv.AppartmentID = appart.ID
	srv.Internet = req.Internet
	srv.Food = req.Food
	srv.LuggageStorage = req.LuggageStorage
	srv.Delivery = req.Delivery
	srv.ClothCleaner = req.ClothCleaner
	model.DB.Create(&srv)

	//!creating service
	appart.Service = srv
	appart.BookStartDate = beginDate
	appart.BookEndDate = endDate
	model.DB.Save(&appart)

	user.Appartments = append(user.Appartments, appart)
	model.DB.Save(&user)
	ctx.JSON(http.StatusOK, gin.H{"message": "Appartmant successfully booked!"})
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
		aprtmt.RentPrice = req.RentPrice
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
