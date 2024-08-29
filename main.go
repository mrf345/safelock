package main

import (
	"embed"
	"fmt"

	desktopEntry "github.com/mrf345/desktop-entry"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS
	//go:embed build/appicon.png
	icon []byte
)

func main() {
	if err := getDesktopEntry().Create(); err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	app := NewApp()
	err := wails.Run(&options.App{
		Title:         "Safelock",
		Width:         410,
		Height:        380,
		DisableResize: true,
		OnStartup:     app.startup,
		OnDomReady:    app.domReady,
		BackgroundColour: &options.RGBA{
			R: 248,
			G: 249,
			B: 250,
			A: 1,
		},
		AssetServer: &assetserver.Options{Assets: assets},
		Bind:        []interface{}{app},
		Linux:       &linux.Options{Icon: icon},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func getDesktopEntry() *desktopEntry.DesktopEntry {
	entry := desktopEntry.New(Name, Version, icon)
	entry.Comment = "Fast & simple drag & drop files encryption tool"
	entry.Categories = "Utility;Security;"
	return entry
}
