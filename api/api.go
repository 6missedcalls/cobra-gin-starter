package api

import (
	"github.com/gin-gonic/gin"

	docs "github.com/6missedcalls/cobra-gin-starter/api/docs"
	routes "github.com/6missedcalls/cobra-gin-starter/api/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	const bp = "/api/v1"
	v1 := router.Group(bp)
	docs.SwaggerInfo.BasePath = bp
	routes.RegisterHealthRoutes(v1)

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
