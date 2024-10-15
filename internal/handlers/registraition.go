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
	fmt.Println("REG_1")
	storage.Players.Create(models.User{Name: data.Name, Role: data.Role})
	fmt.Println("REG_2")

	// Возвращаем JSON-ответ с URL для перенаправления
	c.JSON(http.StatusOK, gin.H{
		"redirect": "/wroom", // URL для перенаправления
		"message":  "Регистрация успешна!",
	})
	fmt.Println("REG_3")

}
