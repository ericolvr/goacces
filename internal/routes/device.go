package routes

import (
	"newaccess/internal/handlers"

	"github.com/gin-gonic/gin"
)

func DeviceRoutes(
	router *gin.Engine,
	deviceHandler *handlers.DeviceHandler,
) {
	routes := router.Group("/api/v1/devices")
	{
		routes.POST("", deviceHandler.Create)
		routes.GET("", deviceHandler.List)
		routes.GET("/:id", deviceHandler.FindByID)
		routes.PUT("/:id", deviceHandler.Update)
		routes.DELETE("/:id", deviceHandler.Delete)
	}
}
