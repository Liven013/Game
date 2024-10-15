package handlers

import (
	"game/internal/models"


	"github.com/gin-gonic/gin"
)

func Wroom(c *gin.Context) {
	var player models.User
	c.BindJSON(&player)
}
