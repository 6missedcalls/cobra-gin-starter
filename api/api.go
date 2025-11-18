package api

import (
	"github.com/gin-gonic/gin"

	routes "github.com/6missedcalls/cobra-gin-starter/api/routes"
	ws "github.com/6missedcalls/cobra-gin-starter/api/websockets"
)

// @BasePath /api/v1
func SetupRouter() *gin.Engine {
	router := gin.Default()

	const bp = "/api/v1"
	v1 := router.Group(bp)
	routes.RegisterSwaggerRoutes(v1)
	routes.RegisterHealthRoutes(v1)
	ws.HelloWebsocketHandler(v1)

	return router
}
