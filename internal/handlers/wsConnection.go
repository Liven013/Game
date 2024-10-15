package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
	}
	defer ws.Close()

	for {
		messegeType, messege, err := ws.ReadMessage()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
			break
		}

		//обработать полученный messege
		if err := ws.WriteMessage(messegeType, messege); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
			break
		}

	}
}
