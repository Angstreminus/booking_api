package main

import (
	"booking/api/controller"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.POST("/login", controller.Login)
}

var router *gin.Engine

func main() {

	router = gin.Default()

	initializeRoutes()

	router.Run()
}
