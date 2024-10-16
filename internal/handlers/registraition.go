package handlers

import (
	"fmt"
	"game/internal/models"
	"game/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormData struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func RegistratePlayer(c *gin.Context) {
	var data FormData

	// Парсим JSON-запрос
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	storage.Players.Create(models.User{Name: data.Name, Role: data.Role})

	broadcast(fmt.Sprintf("New user registered: %s as %s", data.Name, data.Role))
	// Возвращаем JSON-ответ с URL для перенаправления
	c.JSON(http.StatusOK, gin.H{
		"redirect": "/wroom", // URL для перенаправления
		"message":  "Регистрация успешна!",
	})
}
