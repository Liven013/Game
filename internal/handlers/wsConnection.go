package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	Clients  = make(map[*websocket.Conn]bool)
	mu       sync.Mutex
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

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

		mu.Lock()
		Clients[ws] = true
		mu.Unlock()

		RegistratePlayer(c)

		if err := ws.WriteMessage(messegeType, messege); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
			mu.Lock()
			delete(Clients, ws)
			mu.Unlock()
			break
		}

	}
}

func broadcast(message string) {
	mu.Lock()
	defer mu.Unlock()
	for client := range Clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			client.Close()
			delete(Clients, client)
		}
	}
}
