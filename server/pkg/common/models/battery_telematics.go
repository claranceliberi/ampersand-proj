package models

import "gorm.io/gorm"

type BatteryTelematics struct {
	gorm.Model        // adds ID, created_at etc.
	BatteryID    uint // foreign key
	Longitude    string
	Latitude     string
	BatterySoc   float64
	Charging     bool
	ChargingRate float64

	// battery information for the battery telematics
	Battery Battery
}