package server

import (
	"fmt"
	"game/internal/handlers"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

var HostIP net.IP

func StartRouter() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	//router.Static("/static", "./static")

	// Получение IP-адреса
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error retrieving addresses:", err)
		return
	}

	var ip net.IP
	for _, addr := range addrs {
		// Проверяем, что это не адрес loopback
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			ip = ipnet.IP
			break
		}
	}

	if ip == nil {
		fmt.Println("No suitable IP address found")
		return
	}

	fmt.Println("Server is running on:", ip.String())

	SetEndpoints(router)
	handlers.GenerateQRCode(ip.String())
	router.Run(":8080")

}

func SetEndpoints(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

}
