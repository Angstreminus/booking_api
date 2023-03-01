package controller

import (
	model "booking/api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqBody struct {
	AppartmantBook_id uint
}

//count days + additional service
func GetUserSummaryBill(ctx *gin.Context) {
	var (
		inputreqBody  ReqBody
		month_numbers = 0
	)
	if err := ctx.ShouldBindJSON(&inputreqBody); err != nil {
		log.Fatal(err)
	}

	app, er := model.FindAppartmentById(inputreqBody.AppartmantBook_id)
	if er != nil {
		log.Fatal(er)
	}
	months := app.BookEndDate.Sub(app.BookStartDate)

	// few days or less 1 month => price for 1 month
	month_numbers = int(months.Hours() / 24 / 30)
	if month_numbers == 0 {
		month_numbers++
	}
	finalBill := app.Service.Bill() + int32(month_numbers)*int32(app.RentPrice)

	ctx.JSON(http.StatusOK, gin.H{"Your final book bill": finalBill})
}
