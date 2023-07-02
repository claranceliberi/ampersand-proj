package vehicles

import (
	"net/http"
	"time"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddVehicleRequestBody struct {
    Identifier      string    `json:"identifier"`
    ModelName       string    `json:"model"`
    ManufactureDate time.Time   `json:"manufacture_date"`
    Description     string    `json:"description"`
}

// AddVehicle godoc
// @Summary Add a new vehicle
// @Description Add a new vehicle
// @Tags vehicles
// @Accept json
// @Produce json
// @Param body body AddVehicleRequestBody true "body"
// @Router /vehicles [post]
func (h handler) AddVehicle(c *gin.Context) {
	body := AddVehicleRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var vehicle models.Vehicle

	vehicle.Identifier = body.Identifier
	vehicle.ModelName = body.ModelName
	vehicle.ManufactureDate = body.ManufactureDate
	vehicle.Description = body.Description

	if result := h.DB.Create(&vehicle); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &vehicle)
}
