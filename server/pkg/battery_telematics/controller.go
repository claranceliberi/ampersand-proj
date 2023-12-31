package battery_telematics


import (
    "github.com/gin-gonic/gin"

    "gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := r.Group("/battery_telematics")
    routes.POST("/", h.AddBatteryTelematics)
    routes.GET("/", h.GetBatteriesTelematics)
}