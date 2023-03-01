package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const (
	Undefined = ""
	Mansion   = "Mansion"
	House     = "House"
	Room      = "Room"
	Hostel    = "Hostel"
)

func GetHouseType(house string) (string, error) {
	switch house {
	case Undefined:
		return Undefined, nil
	case House:
		return House, nil
	case Room:
		return Room, nil
	case Hostel:
		return Hostel, nil
	}

	return Undefined, errors.New("Unknown house type: " + house)
}

type Appartment struct {
	gorm.Model
	UserRefer     uint
	HousingType   string
	RentPrice     int
	BookStartDate time.Time
	BookEndDate   time.Time
	Service       Service
}

//!Searching Appartment
func FindAppartmentById(id uint) (appratment Appartment, err error) {
	if err := DB.Where("id = ?", id).First(&appratment).Error; err != nil {
		return appratment, err
	}
	return appratment, nil
}
