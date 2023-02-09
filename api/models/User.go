package model

type User struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	Name        string
	Password    string
	Appartments []Appartment
}
