package handlers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func GetHostIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error retrieving addresses:", err)
		return ""
	}

	var ip net.IP
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP
				break
			}
		}
	}

	if ip == nil {
		fmt.Println("No suitable IP address found")
		return ""
	}
	fmt.Println("Server is running on:", ip.String())
	return ip.String()
}

func QRGenerator(c *gin.Context) {
	qrData := fmt.Sprintf("http://%s:8080", GetHostIP())

	png, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate QR code")
		return
	}

	c.Data(http.StatusOK, "image/png", png)
}

func Start(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}
