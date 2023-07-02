package models

import "gorm.io/gorm"


type Agent struct {
    gorm.Model  // adds ID, created_at etc.
    FirstName  string    `gorm:"type:varchar"`
    LastName   string    `gorm:"type:varchar"`
    PhoneNumber string   `gorm:"type:varchar"`
    Email      string   `gorm:"type:varchar"`

	  // swaps is a one-to-many relationship, i.e. a agent can have many swaps
    Swaps []Swap
}
