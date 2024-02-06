package model

import (
	"time"
	// govalidador
)

type Base struct {
	ID        string    `json:"valid:uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
