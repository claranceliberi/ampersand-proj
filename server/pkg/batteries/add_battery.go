package batteries

import (
	"net/http"
	"time"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddBatteryRequestBody struct {
	Identifier      string    `json:"identifier"`
	ModelName       string    `json:"model_name"`
	LastSeenOnline  time.Time `json:"last_seen_online"`
	ManufactureDate time.Time `json:"manufacture_date"`
}

// AddBattery godoc
// @Summary Add a new battery
// @Description Add a new battery
// @Tags batteries
// @Accept json
// @Produce json
// @Param body body AddBatteryRequestBody true "body"
// @Router /batteries [post]
func (h handler) AddBattery(c *gin.Context) {
	body := AddBatteryRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var battery models.Battery

	battery.Identifier = body.Identifier
	battery.ModelName = body.ModelName
	battery.LastSeenOnline = body.LastSeenOnline
	battery.ManufactureDate = body.ManufactureDate



	if result := h.DB.Create(&battery); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &battery)
}
