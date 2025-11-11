package api

import (
	routes "github.com/6missedcalls/cobra-gin-starter/api/routes"
	"github.com/gin-gonic/gin"
	docs "github.com/6missedcalls/cobra-gin-starter/api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"
	routes.RegisterHealthRoutes(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}

