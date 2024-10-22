package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Wroom(c *gin.Context) {
	// Отправляем HTML-страницу
	c.HTML(http.StatusOK, "w_page.html", gin.H{"messege": "new wroom"})
}
