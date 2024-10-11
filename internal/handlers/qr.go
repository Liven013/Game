package handlers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

var HostIP net.IP

func GetHostIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error retrieving addresses:", err)
		return ""
	}

	var ip net.IP
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// Проверяем, что это IPv4-адрес
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
	HostIP = ip
	fmt.Println("Server is running on:", ip.String())
	return ip.String()
}

func QRGenerator(c *gin.Context) {
	qrData := fmt.Sprintf("http://%s:8080", GetHostIP())

	// Генерация QR-кода в памяти и отправка изображения PNG
	png, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate QR code")
		return
	}

	// Отправляем изображение PNG в ответе
	c.Data(http.StatusOK, "image/png", png)
}

func Start(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func GenerateQRCode(ip string) error {
	qrFile := "qrcode.png"
	err := qrcode.WriteFile(ip, qrcode.Medium, 256, qrFile)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return err
	}
	fmt.Println("QR code generated and saved as:", qrFile)
	return nil
}
