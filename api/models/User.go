package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	Name        string
	Password    string
	Appartments []Appartment
}

func Encrypt(passwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
}

func Verify(hasedPassword, passwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(passwd))
}
