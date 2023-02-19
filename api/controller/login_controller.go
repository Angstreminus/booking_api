package controller

import (
	"booking/api/auth"
	model "booking/api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	usr := new(model.User)

	if err := c.BindJSON(usr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := SingIn(usr.Email, usr.Password)

	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func SingIn(email, password string) (string, error) {
	var (
		err error
		DB  *gorm.DB
	)
	user := model.User{}

	err = DB.Debug().Model(model.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = model.Verify(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(uint32(user.ID))
}
