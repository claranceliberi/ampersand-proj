package drivers

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddDriverRequestBody struct {
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    PhoneNumber string `json:"phone_number"`
    Email       string `json:"email"`
}

// AddDriver godoc
// @Summary Add a new driver
// @Description Add a new driver
// @Tags drivers
// @Accept json
// @Produce json
// @Param body body AddDriverRequestBody true "body"
// @Router /drivers [post]
func (h handler) AddDriver(c *gin.Context) {
    body := AddDriverRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

	var driver models.Driver

	driver.FirstName = body.FirstName
	driver.LastName = body.LastName
	driver.PhoneNumber = body.PhoneNumber
	driver.Email = body.Email


    if result := h.DB.Create(&driver); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &driver)
}