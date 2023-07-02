package agents

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

    routes := r.Group("/agents")
    routes.POST("/", h.AddAgent)
    routes.GET("/", h.GetAgents)
}