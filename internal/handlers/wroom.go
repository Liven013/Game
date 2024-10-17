package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Wroom(c *gin.Context) {
	// Отправляем HTML-страницу без выполнения WSConnection
	c.HTML(http.StatusOK, "w_page.html", gin.H{"messege": "new wroom"})
}
