package main

import (
	"booking/api/controller"
	model "booking/api/models"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.POST("/appartments/new", controller.CreateAppartment)
	router.POST("/user/bill", controller.BookHousing)
	router.GET("/user/bill", controller.GetUserSummaryBill)
}

var router *gin.Engine

func main() {
	model.InitDB()

	router = gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	initializeRoutes()

	router.Run()
}
