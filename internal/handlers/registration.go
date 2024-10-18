package handlers

import (
	"fmt"
	"game/internal/models"
	"game/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type FormData struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// HTTP-обработчик для регистрации пользователя
func RegistratePlayer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"redirect": fmt.Sprint("http://" + HostIP + ":8080/wroom")})
}

// WebSocket-обработчик
func WSConnection(c *gin.Context) {
	var data FormData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}
	defer ws.Close()

	storage.Users.Create(models.User{Name: data.Name, Role: data.Role}, ws)
	storage.Users.Broadcast(fmt.Sprintf("New user registered: %s as %s", data.Name, data.Role))

	go func() {
		for {
			storage.Users.SendAllUsers()
		}
	}()
	select {}
}
