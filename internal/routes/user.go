package routes

import (
	"newaccess/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
) {
	routes := router.Group("/api/v1/users")
	{
		routes.POST("", userHandler.Create)
		routes.GET("", userHandler.List)
		routes.GET("/:id", userHandler.FindByID)
		routes.PUT("/:id", userHandler.Update)
		routes.DELETE("/:id", userHandler.Delete)
	}
}
