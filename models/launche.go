package models

import "time"

type Launche struct {
	Id        int       `json:"id" db:"id"`
	Provider  bool      `json:"provider" db:"provider"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Launches []Launche
