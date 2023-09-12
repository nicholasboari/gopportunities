package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/nicholasboari/gopportunities/docs"
	"github.com/nicholasboari/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initilizeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	v1.GET("/opening", handler.FindOpeningHandler)
	v1.POST("/opening", handler.CreateOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)
	v1.GET("/openings", handler.ListOpeningsHandler)

	// initialize swagger
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
