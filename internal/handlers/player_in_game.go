package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var Clients = make(map[chan string]bool)

func SseHandler(c *gin.Context) {
	sse(c)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			message := fmt.Sprintf("Current time: %s", time.Now().Format(time.RFC1123))
			broadcast(message)
		}
	}()
}
func sse(c *gin.Context) {
	// Устанавливаем заголовки для SSE
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	// Создаем канал для отправки сообщений клиенту
	messageChan := make(chan string)
	Clients[messageChan] = true

	// Отправляем сообщения клиенту
	go func() {
		for {
			msg := <-messageChan
			c.SSEvent("message", msg) // Отправляем событие
		}
	}()

	// Удаляем клиентский канал при закрытии соединения
	defer func() {
		delete(Clients, messageChan)
	}()

	// Блокируем горутину
	<-c.Request.Context().Done()
}

func broadcast(message string) {
	for client := range Clients {
		client <- message // Отправляем сообщение всем подключенным клиентам
	}
}
