package server

import (
	"fmt"
	"game/internal/handlers"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.LoadHTMLGlob("front_part/templates/*")

	router.Static("front_part", "./front_part")

	SetEndpoints(router)

	// Получение IP-адреса хоста перед тем, как открыть браузер
	openBrowser(fmt.Sprintf("http://%s:8080/qr", handlers.HostIP))

	// Запуск сервера
	router.Run(":8080")
}

func SetEndpoints(g *gin.Engine) {
	g.GET("/", handlers.Start)
	g.GET("/qr", handlers.QRGenerator)

	g.POST("/submit", handlers.RegistratePlayer)

	g.GET("/ws", handlers.WSConnection)
	g.GET("/wroom", handlers.Wroom)

	//управление Каталогом игроков
	g.GET("/players", handlers.GetAll)
	g.POST("/players", handlers.CreateUser)
	g.GET("/players/:id", handlers.GetOne)
	g.DELETE("/players/:id", handlers.DelByID)

}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		fmt.Println("Unsupported platform")
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to open browser:", err)
	}
}
