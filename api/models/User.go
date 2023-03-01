package model

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role        string
	Email       string
	Password    string
	Appartments []Appartment `gorm:"foreignKey:UserRefer"`
	Token       string
}

type Jwt struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	*jwt.StandardClaims
}

func Encrypt(passwd string) []byte {
	pswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return pswd
}

func CheckAdmin(admin_ID string) bool {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	return admin_ID == os.Getenv("ADMIN_ROLE_ID")
}

func Verify(hasedPassword, passwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(passwd))
}
func SetRole(role string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error occured: ", err)
	}

	if role == os.Getenv("ADMIN_ROLE_ID") {
		return "Admin"
	} else {
		return "User"
	}
}
