package agents

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetAgents godoc
// @Summary Get all agents
// @Description Get all agents
// @Tags agents
// @Accept json
// @Produce json
// @Router /agents [get]
func (h handler) GetAgents(c *gin.Context) {
    var agents []models.Agent

    if result := h.DB.Find(&agents); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &agents)
}