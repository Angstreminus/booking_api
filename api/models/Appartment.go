package model

import "time"

type Appartment struct {
	ID        int
	Type      string
	BookStart time.Time
	BookEnd   time.Time
	Service   Service
}
