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

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{}
	u.Email = input.Email
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
