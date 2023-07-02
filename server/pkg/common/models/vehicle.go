package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model             // adds ID, created_at etc.
	Identifier      string `gorm:"type:varchar;unique;"`
	ModelName       string
	ManufactureDate time.Time `gorm:"type:date"`
	Description     string    `gorm:"type:varchar"`

	// swaps is a one-to-many relationship, i.e. a vehicle can have many swaps
	Swaps []Swap
}
