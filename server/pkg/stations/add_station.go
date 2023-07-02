package stations

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddStationsRequestBody struct {
	Location    string `json:"location"`
	Description string `json:"description"`
}

// AddStation godoc
// @Summary Add a station
// @Description Add a station
// @Tags stations
// @Accept json
// @Produce json
// @Param body body AddStationsRequestBody true "body"
// @Router /stations [post]
func (h handler) AddStation(c *gin.Context) {
	body := AddStationsRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var stations models.Station

	stations.Location = body.Location
	stations.Description = body.Description

	if result := h.DB.Create(&stations); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &stations)
}
