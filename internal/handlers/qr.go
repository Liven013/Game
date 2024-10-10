package handlers

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(ip string) {
	qrFile := "qrcode.png"
	err := qrcode.WriteFile(fmt.Sprintf("http://%s:8080", ip), qrcode.Medium, 256, qrFile)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}
	fmt.Println("QR code generated and saved as:", qrFile)

	var cmd *exec.Cmd
	switch os := runtime.GOOS; os {
	case "darwin": // macOS
		cmd = exec.Command("open", qrFile)
	case "windows": // Windows
		cmd = exec.Command("cmd", "/c", "start", qrFile)
	case "linux": // Linux
		cmd = exec.Command("xdg-open", qrFile)
	default:
		fmt.Println("Unsupported OS:", os)
		return
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error opening QR code:", err)
	}
}
