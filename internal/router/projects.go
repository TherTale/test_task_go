package router

import (
	"contact-center-system/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine, projectHandler *handlers.ProjectHandler) {
	projectGroup := router.Group("/projects")
	{
		projectGroup.POST("/", projectHandler.CreateProject)
		projectGroup.GET("/", projectHandler.GetProject)

		projectGroup.DELETE("/:id", projectHandler.DeleteProject)
		projectGroup.PUT("/:id", projectHandler.PutProject)
		// Add more project routes as needed

		projectGroup.POST("/:id/operator/:operatorId", projectHandler.AddOperatorProject)
		projectGroup.DELETE("/:id/operator/:operatorId", projectHandler.DeleteOperatorProject)
	}
}
