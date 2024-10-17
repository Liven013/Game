package handlers

import (
	"fmt"
	"game/internal/models"
	"game/internal/storage"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type FormData struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

var (
	Clients  = make(map[*websocket.Conn]bool)
	mu       sync.Mutex
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// HTTP-обработчик для регистрации пользователя
func RegistratePlayer(c *gin.Context) {
	var data FormData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	storage.Players.Create(models.User{Name: data.Name, Role: data.Role})

	broadcast(fmt.Sprintf("New user registered: %s as %s", data.Name, data.Role))
	c.JSON(http.StatusOK, gin.H{"redirect": fmt.Sprint("http://" + HostIP + ":8080/wroom")})

}

// WebSocket-обработчик
func WSConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}
	defer ws.Close()

	mu.Lock()
	Clients[ws] = true
	mu.Unlock()
	/*
		for {
			// Читаем сообщения из WebSocket
			_, message, err := ws.ReadMessage()
			if err != nil {
				mu.Lock()
				delete(Clients, ws)
				mu.Unlock()
				break
			}
		}
	*/
}

// Функция для рассылки сообщений всем подключённым клиентам
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
