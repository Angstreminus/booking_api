//!!TODO metki

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
	BookStartDate time.Time
	BookEndDate   time.Time
	Service       Service
}
