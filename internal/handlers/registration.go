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
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	// defer ws.Close()
	var data FormData

	if err := ws.ReadJSON(&data); err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
	}
	fmt.Println("уже тут2")
	storage.Users.Create(models.User{Name: data.Name, Role: data.Role}, ws)
	fmt.Println("уже тут1")
	storage.Users.Broadcast(fmt.Sprintf("New user registered: %s as %s", data.Name, data.Role))
	fmt.Println("уже тут")
	storage.Users.SendAllUsers()
	fmt.Println("сюдааа")
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			// Закрываем соединение, если возникает ошибка чтения
			fmt.Println("Ошибка чтения сообщения:", err)
			break
		}
	}
	fmt.Println("и че такое?")
	storage.Users.Delete(ws)
	storage.Users.SendAllUsers()
}
