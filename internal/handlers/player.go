package handlers

import (
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

	CreatePayer()

	// Возвращаем JSON-ответ с URL для перенаправления
	c.JSON(http.StatusOK, gin.H{
		"redirect": "/waiting", // URL для перенаправления
		"message":  "Регистрация успешна!",
	})
}
