
package vehicles

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetVehicles godoc
// @Summary Get all vehicles
// @Description Get all vehicles
// @Tags vehicles
// @Accept json
// @Produce json
// @Router /vehicles [get]
func (h handler) GetVehicles(c *gin.Context) {
    var vehicles []models.Vehicle

    if result := h.DB.Find(&vehicles); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &vehicles)
}