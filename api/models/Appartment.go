package model

import (
	"errors"
	"time"
)

type Housing struct {
	houseType string
}

func (h *Housing) String() string {
	return h.houseType
}

var (
	Undefined = Housing{""}
	Mansion   = Housing{"Mansion"}
	House     = Housing{"House"}
	Room      = Housing{"Room"}
	Hostel    = Housing{"Hostel"}
)

func GetHouseType(house string) (Housing, error) {
	switch house {
	case Mansion.houseType:
		return Mansion, nil
	case House.houseType:
		return House, nil
	case Room.houseType:
		return Room, nil
	case Hostel.houseType:
		return Hostel, nil
	}

	return Undefined, errors.New("Unknown house type: " + house)
}

type Appartment struct {
	ID        int
	Type      Housing
	BookStart time.Time
	BookEnd   time.Time
	Service   Service
}
