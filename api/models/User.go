package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string
	Password    string
	Appartments []Appartment `gorm:"foreignKey:UserRefer"`
}

func Encrypt(passwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
}

func Verify(hasedPassword, passwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(passwd))
}

func (u *User) SaveUser() (*User, error) {

	err := DB.Create(&u).Error

	if err != nil {
		log.Fatal(err)
		return &User{}, err
	}
	return u, err
}
