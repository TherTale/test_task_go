package router

import (
	"contact-center-system/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine) {
	projectGroup := router.Group("/projects")
	{
		projectGroup.GET("/", handlers.GetProject)
		projectGroup.POST("/", handlers.CreateProject)

		projectGroup.DELETE("/", handlers.DeleteProject)
		projectGroup.PUT("/", handlers.PutProject)
		// Add more project routes as needed
	}
}
