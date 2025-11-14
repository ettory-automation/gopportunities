package router

import (
	"github.com/ettory-automation/gopportunities/handler"
	"github.com/gin-gonic/gin"

	"github.com/ettory-automation/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"             // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"         // gin-swagger middleware
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath 
	v1 := router.Group(basePath)
	{
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // Correção: r para router
}
