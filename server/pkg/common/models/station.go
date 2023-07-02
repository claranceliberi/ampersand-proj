package models

import "gorm.io/gorm"

type Station struct {
	gorm.Model
	Location    string `gorm:"type:varchar;"`
	Description string `gorm:"type:varchar;"`

	  // One-To-Many relationship (has many) with Swap, use StationID as foreign key
	  Swaps []Swap 
}
