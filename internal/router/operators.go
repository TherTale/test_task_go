package router

import (
	"contact-center-system/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterOperatorRoutes(router *gin.Engine, operatorHandler *handlers.OperatorHandler) {
	operatorGroup := router.Group("/operators")
	{
		operatorGroup.GET("/", operatorHandler.GetOperators)
		operatorGroup.POST("/", operatorHandler.CreateOperators)

		operatorGroup.DELETE("/:id", operatorHandler.DeleteOperators)
		operatorGroup.PUT("/:id", operatorHandler.PutOperators)
		// Add more operator routes as needed
	}
}
