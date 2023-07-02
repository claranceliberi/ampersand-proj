package batteries

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetBatteries godoc
// @Summary Get all batteries
// @Description Get all batteries
// @Tags batteries
// @Accept json
// @Produce json
// @Router /batteries [get]
func (h handler) GetBatteries(c *gin.Context) {
	var batteries []models.Battery

	if result := h.DB.Find(&batteries); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &batteries)
}
