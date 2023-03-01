package controller

import (
	model "booking/api/models"
	response "booking/api/responses"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   string `json:"role_id" binding:"required"`
}

type ReqLoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResLoginBody struct {
	Token string `json:"token"`
}

func NewWithClaims(claims jwt.Claims) (ss string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString([]byte("codedoct"))
	return
}

func UpdateToken(user model.User, ss string) error {
	user.Token = ss
	return model.DB.Save(&user).Error
}

func findUsr(email string) (user model.User, err error) {
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := model.User{
		Role:        model.SetRole(input.RoleID),
		Email:       input.Email,
		Password:    string(model.Encrypt(input.Password)),
		Appartments: nil,
	}
	err := model.DB.Create(&u).Error
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func LoginCheck(loginpt *ReqLoginBody) (*ResLoginBody, int, error) {

	var (
		resBody ResLoginBody
		usr     = model.User{}
	)

	user, err := findUsr(loginpt.Email)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("UserNotFound")
	}

	if err = model.Verify(user.Password, loginpt.Password); err != nil {
		return nil, http.StatusBadRequest, errors.New("PASSWORD INCORRECT")
	}
	claims := model.Jwt{
		ID:    int(user.ID),
		Email: user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	ss, err := NewWithClaims(claims)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	resBody.Token = ss
	model.DB.Model(&usr).Where("email = ?", loginpt.Email).Update("token", ss)

	return &resBody, http.StatusOK, nil
}
func Login(ctx *gin.Context) {
	var (
		reqBody ReqLoginBody
	)
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	resBody, errStatus, err := LoginCheck(&reqBody)
	if err != nil {
		response.Error(ctx, errStatus, err.Error())
		return
	}

	response.Json(ctx, http.StatusOK, resBody)

}
