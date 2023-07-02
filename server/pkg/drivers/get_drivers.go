package drivers

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetDrivers godoc
// @Summary Get all drivers
// @Description Get all drivers
// @Tags drivers
// @Accept json
// @Produce json
// @Router /drivers [get]
func (h handler) GetDrivers(c *gin.Context) {
    var driver []models.Driver

    if result := h.DB.Find(&driver); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &driver)
}