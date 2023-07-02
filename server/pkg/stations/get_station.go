package stations

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetStations godoc
// @Summary Get all stations
// @Description Get all stations
// @Tags stations
// @Accept json
// @Produce json
// @Router /stations [get]
func (h handler) GetStations(c *gin.Context) {
    var stations []models.Station

    if result := h.DB.Find(&stations); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &stations)
}