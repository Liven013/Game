package handlers

import (
	"fmt"
	"game/internal/models"
	"game/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	users := storage.Players.GetAll()
	c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	//исправить на ключ по имени! и обрабатывать совпадения
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}
	storage.Players.Create(user)
	GetAll(c)
}

func GetOne(c *gin.Context) {
	id := c.Param("id")
	user, err := storage.Players.GetOne(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
	}
	c.IndentedJSON(http.StatusOK, user)
}

func DelByID(c *gin.Context) {
	id := c.Param("id")
	err := storage.Players.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "delete user"})
}
