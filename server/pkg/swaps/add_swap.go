package swaps

import (
	"errors"
	"log"
	"math"
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/claranceliberi/ampersand/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

type AddSwapRequestBody struct {
	BatteryInID  float64 `json:"battery_in"`
	BatteryOutID float64 `json:"battery_out"`
	DriverID     float64 `json:"driver"`
	StationID    float64 `json:"station"`
	AgentID      float64 `json:"agent"`
	VehicleID    float64 `json:"vehicle"`
}

const COST_PER_PERCENT = 12.5

// AddSwap godoc
// @Summary Add a swap
// @Description Add a swap
// @Tags swaps
// @Accept json
// @Produce json
// @Param body body AddSwapRequestBody true "body"
// @Router /swaps [post]
func (h handler) AddSwap(c *gin.Context) {
	body := AddSwapRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var batteryIn models.Battery
	existsIn, err := utils.RecordExistsByID(h.DB, &batteryIn, body.BatteryInID)

	// Handle the error and existence check with utility function
	utils.HandleDatabaseCheck(c, existsIn, err, "No battery was found with ID: "+utils.ToString(body.BatteryInID))

	var batteryOut models.Battery
	existsOut, err := utils.RecordExistsByID(h.DB, &batteryOut, body.BatteryInID)

	// Handle the error and existence check with utility function
	utils.HandleDatabaseCheck(c, existsOut, err, "No battery was found with ID: "+utils.ToString(body.BatteryOutID))

	var driver models.Driver
	existsDriver, err := utils.RecordExistsByID(h.DB, &driver, body.DriverID)

	// Handle the error and existence check with utility function
	utils.HandleDatabaseCheck(c, existsDriver, err, "No Driver was found with ID: "+utils.ToString(body.DriverID))

	var station models.Driver
	existsStation, err := utils.RecordExistsByID(h.DB, &station, body.StationID)

	// Handle the error and existence check with utility function
	utils.HandleDatabaseCheck(c, existsStation, err, "No Station was found with ID: "+utils.ToString(body.StationID))

	var agent models.Agent
	existsAgent, err := utils.RecordExistsByID(h.DB, &agent, body.AgentID)

	// Handle the error and existence check with utility function
	utils.HandleDatabaseCheck(c, existsAgent, err, "No Agent was found with ID: "+utils.ToString(body.AgentID))

	var vehicle models.Agent
	existsVehicle, err := utils.RecordExistsByID(h.DB, &vehicle, body.VehicleID)

	// Handle the error and existence check with utility function
	utils.HandleDatabaseCheck(c, existsVehicle, err, "No Vehicle was found with ID: "+utils.ToString(body.VehicleID))

	// Get latest battery soc
	batteryInSoc, err := h.getLatestBatterySoc(body.BatteryInID)
	if err != nil {
		// Handle error
		log.Printf("Error getting BatteryIn Soc: %v", err)
		return
	}

	// Get latest battery soc
	batteryOutSoc, err := h.getLatestBatterySoc(body.BatteryOutID)
	if err != nil {
		// Handle error
		log.Printf("Error getting BatteryOut Soc: %v", err)
		return
	}

	// Calculate cost and save it to swap model
	cost := 0.0
	if batteryInSoc > batteryOutSoc { // if batteryInSoc is greater than batteryOutSoc
		cost = math.Ceil((batteryInSoc - batteryOutSoc) * COST_PER_PERCENT)
	} else {
		err := errors.New("Can not replace a battery with a lower soc")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var swap models.Swap

	swap.BatteryInID = uint(body.BatteryInID)
	swap.BatteryInSoc = batteryInSoc
	swap.BatteryOutID = uint(body.BatteryOutID)
	swap.BatteryOutSoc = batteryOutSoc
	swap.Cost = cost
	swap.DriverID = uint(body.DriverID)
	swap.StationID = uint(body.StationID)
	swap.AgentID = uint(body.AgentID)
	swap.VehicleID = uint(body.VehicleID)

	if result := h.DB.Create(&swap); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusCreated, &swap)
}

// the function to get the latest BatterySoc for a Battery
func (h handler) getLatestBatterySoc(batteryId float64) (float64, error) {
	var batteryTelematics models.BatteryTelematics
	// find the latest record for the specified battery
	if result := h.DB.Where("battery_id = ?", uint(batteryId)).Order("created_at desc").First(&batteryTelematics); result.Error != nil {
		return 0, result.Error
	}
	return batteryTelematics.BatterySoc, nil
}
