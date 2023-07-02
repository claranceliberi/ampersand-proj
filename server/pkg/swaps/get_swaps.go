package swaps

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetSwaps godoc
// @Summary Get all swaps
// @Description Get all swaps
// @Tags swaps
// @Accept json
// @Produce json
// @Router /swaps [get]
func (h handler) GetSwaps(c *gin.Context) {
	var swaps []models.Swap

	if result := h.DB.Preload("BatteryIn").
		Preload("BatteryOut").
		Preload("Driver").
		Preload("Station").
		Preload("Agent").
		Preload("Vehicle").
		Order("created_at desc").
		Find(&swaps); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &swaps)
}
