package swaps

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

    routes := r.Group("/swaps")
    routes.POST("/", h.AddSwap)
    routes.GET("/", h.GetSwaps)
    routes.GET("/:id", h.GetSwap)
}