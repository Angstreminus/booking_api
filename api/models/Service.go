package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	AppartmentID   uint
	Internet       bool
	Food           bool
	LuggageStorage bool
	Delivery       bool
	ClothCleaner   bool
}

func (s *Service) Bill() int32 {
	sum := 0

	if s.Internet {
		sum += 100
	}

	if s.Food {
		sum += 250
	}

	if s.LuggageStorage {
		sum += 300
	}

	if s.Delivery {
		sum += 100
	}

	if s.ClothCleaner {
		sum += 500
	}
	return int32(sum)
}
