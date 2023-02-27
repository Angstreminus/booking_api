package main

import (
	"booking/api/controller"
	model "booking/api/models"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}

var router *gin.Engine

func main() {
	model.InitDB()

	router = gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	initializeRoutes()

	router.Run()
}
