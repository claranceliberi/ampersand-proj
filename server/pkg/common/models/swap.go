package models

import "gorm.io/gorm"

type Swap struct {
	gorm.Model           // adds ID, created_at etc.
	BatteryInID     uint // foreign key
	BatteryInSoc    float64
	BatteryOutID    uint // foreign key
	BatteryOutSoc   float64
	Cost            float64
	DistanceCovered float64 // in km
	EnergyConsumed  float64 // in kWh
	DriverID        uint    // foreign key
	StationID       uint    // foreign key
	AgentID         uint    // foreign key
	VehicleID       uint    // foreign key

	// swapped in battery
	BatteryIn Battery
	// swapped out battery
	BatteryOut Battery
	// driver who performed the swap
	Driver Driver
	// station where the swap took place
	Station Station
	// agent who performed the swap
	Agent Agent
	// vehicle which was swapped
	Vehicle Vehicle
}
