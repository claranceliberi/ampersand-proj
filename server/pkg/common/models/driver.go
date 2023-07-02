package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model  // adds ID, created_at etc.
	FirstName   string
	LastName    string
	PhoneNumber string `gorm:"type:varchar"`
	Email       string `gorm:"type:varchar"`

	// swaps is a one-to-many relationship, i.e. a driver can have many swaps
	Swaps []Swap 
}
