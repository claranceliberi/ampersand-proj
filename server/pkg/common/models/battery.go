package models

import (
	"time"
	"gorm.io/gorm"
)



type Battery struct {
	gorm.Model  // adds ID, created_at etc.
    Identifier       string    `gorm:"type:varchar;unique;"`
    ModelName            string
    LastSeenOnline   time.Time `gorm:"type:timestamp without time zone;"`
    ManufactureDate  time.Time `gorm:"type:date"`

	 // swaps is a one-to-many relationship, i.e. a battery can have many swaps
	 // SwapsIn is the list of swaps where this battery was swapped in
	 SwapsIn  []Swap 	`gorm:"foreignKey:BatteryInID"`
	 // SwapsOut is the list of swaps where this battery was swapped out
	 SwapsOut []Swap 	`gorm:"foreignKey:BatteryOutID"`

	 // batteryTelematics
	 BatteryTelematics []BatteryTelematics 		`gorm:"foreignKey:BatteryID"`
}
