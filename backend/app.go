package backend

import (
	"embed"
	"runtime"

	desktopEntry "github.com/mrf345/desktop-entry"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

func NewApp(icon []byte, assets embed.FS) (*App, *options.App) {
	height := 380
	isWindows := runtime.GOOS == "windows"

	if isWindows {
		height += 30
	}

	app := &App{}
	return app, &options.App{
		Title:         Name,
		Width:         410,
		Height:        height,
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
	}
}

func NewDesktopEntry(icon []byte) *desktopEntry.DesktopEntry {
	entry := desktopEntry.New(Name, Version, icon)
	entry.Comment = "Fast & simple drag & drop files encryption tool"
	entry.Categories = "Utility;Security;"
	return entry
}
