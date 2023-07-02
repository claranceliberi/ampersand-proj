package swaps

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetSwaps godoc
// @Summary Get Swap
// @Description Get Swap
// @Tags swaps
// @Accept json
// @Produce json
// @Param id path string true "ID of the swap"
// @Router /swaps/{id} [get]
func (h handler) GetSwap(c *gin.Context) {
	id := c.Param("id")

	var swap models.Swap

	if result := h.DB.Preload("BatteryIn").Preload("BatteryOut").Preload("Driver").Preload("Station").Preload("Agent").Preload("Vehicle").First(&swap, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &swap)
}
