package main

import (
	"embed"
	"fmt"

	"github.com/mrf345/safelock/backend"
	"github.com/wailsapp/wails/v2"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS
	//go:embed build/appicon.png
	icon []byte
)

func main() {
	if err := backend.NewDesktopEntry(icon).Create(); err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	_, config := backend.NewApp(icon, assets)

	if err := wails.Run(config); err != nil {
		println("Error:", err.Error())
	}
}
