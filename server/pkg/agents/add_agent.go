package agents

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddAgentRequestBody struct {
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    PhoneNumber string `json:"phone_number"`
    Email       string `json:"email"`
}



// AddAgent godoc
// @Summary Add a new agent
// @Description Add a new agent
// @Tags agents
// @Accept json
// @Produce json
// @Param body body AddAgentRequestBody true "body"
// @Router /agents [post]
func (h handler) AddAgent(c *gin.Context) {
    body := AddAgentRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var agent models.Agent

    agent.FirstName = body.FirstName
    agent.LastName = body.LastName
    agent.PhoneNumber = body.PhoneNumber
    agent.Email = body.Email



    if result := h.DB.Create(&agent); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &agent)
}