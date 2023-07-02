package db

import (
	"log"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	// open a db connection
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// create the tables if they doesn't exist
	error := db.AutoMigrate(&models.Vehicle{}, &models.Battery{}, &models.Station{}, &models.Driver{}, &models.Agent{}, &models.Swap{}, &models.BatteryTelematics{})

	if error != nil {
		log.Fatalln(error)
	}

	return db
}
