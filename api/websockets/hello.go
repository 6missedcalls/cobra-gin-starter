package websockets

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

// HelloWebsocketHandler godoc
// @Summary Basic WebSocket Implementation
// @Schemes
// @Description A simple WebSocket endpoint that sends a greeting message.
// @Tags example
// @Accept json
// @Produce json
// @Success 101 {string} Hello, WebSocket!
// @Router /ws [get]
func HelloWebsocketHandler(r *gin.RouterGroup) {
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to upgrade to websocket: %v", err)
			return
		}
		defer func() {
			if cerr := conn.Close(); cerr != nil {
				log.Printf("Error closing websocket connection: %v", cerr)
			}
		}()

		for {
			mType, data, err := conn.ReadMessage()
			if err != nil {
				break
			}
			message := "Hello, WebSocket!, you said: " + string(data)
			if err := conn.WriteMessage(mType, []byte(message)); err != nil {
				break
			}
		}
	})
}
