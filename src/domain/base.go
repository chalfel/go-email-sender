package domain

import "time"

// Base is the base struct that all others structs needs to implements
type Base struct {
	ID        int
	CreatedAt time.Timer
	UpdatedAt time.Timer
}
