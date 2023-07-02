package battery_telematics

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetBatteriesTelematics godoc
// @Summary Get all batteries telematics
// @Description Get all batteries telematics
// @Tags battery_telematics
// @Accept json
// @Produce json
// @Router /battery_telematics [get]
func (h handler) GetBatteriesTelematics(c *gin.Context) {
    var batteries []models.BatteryTelematics

    if result := h.DB.Find(&batteries); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &batteries)
}