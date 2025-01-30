package server

import (
	"github.com/MrD0511/deck/deck-cli/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/pods", handlers.GetPodsHandler)
	}
}


