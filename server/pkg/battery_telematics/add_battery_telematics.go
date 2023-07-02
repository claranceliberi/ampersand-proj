package battery_telematics

import (
	"log"
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/claranceliberi/ampersand/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

type AddBatteryTelematicsRequestBodyEntity struct {
	BatteryId    float64 `json:"battery_id"`
	Longitude    string  `json:"longitude"`
	Latitude     string  `json:"latitude"`
	BatterySoc   float64 `json:"battery_soc"`
	Charging     bool    `json:"charging"`
	ChargingRate float64 `json:"charging_rate"`
}

// AddBatteryTelematics godoc
// @Summary Add a battery telematics
// @Description Add a battery telematics
// @Tags battery_telematics
// @Accept json
// @Produce json
// @Param body body AddBatteryTelematicsRequestBodyEntity true "body"
// @Router /battery_telematics [post]
func (h handler) AddBatteryTelematics(c *gin.Context) {
	body := AddBatteryTelematicsRequestBodyEntity{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var battery models.Battery
	exists, err := utils.RecordExistsByID(h.DB, &battery, body.BatteryId)

	// Handle the error and existence check with utility function
	utils.HandleDatabaseCheck(c, exists, err, "No battery was found with ID: "+utils.ToString(body.BatteryId))

	if !exists {
		return
	}

	go func() {
		var batteryTelematics models.BatteryTelematics
		batteryTelematics.BatteryID = uint(body.BatteryId)
		batteryTelematics.Longitude = body.Longitude
		batteryTelematics.Latitude = body.Latitude
		batteryTelematics.BatterySoc = body.BatterySoc
		batteryTelematics.Charging = body.Charging
		batteryTelematics.ChargingRate = body.ChargingRate

		if result := h.DB.Create(&batteryTelematics); result.Error != nil {
			log.Printf("Error in creating battery telematics: %v", result.Error)
			return
		}
	}()

	c.JSON(http.StatusCreated, gin.H{"status": "The request has been accepted and is being processed"})
}
