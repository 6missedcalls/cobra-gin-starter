package routes

import (
	docs "github.com/6missedcalls/cobra-gin-starter/api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func RegisterSwaggerRoutes(rg *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
