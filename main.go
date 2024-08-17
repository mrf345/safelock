//go:generate fyne bundle -o bundled.go assets

package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"github.com/mrf345/safelock/gui"
)

func main() {
	app := app.NewWithID("safelock")
	app.Settings().SetTheme(gui.GetCustomTheme(false))
	window := app.NewWindow("Home")
	window.SetFixedSize(true)
	window.SetTitle(fmt.Sprintf("Safelock v%s", gui.Version))
	window.Resize(gui.DefaultSize)
	window.SetContent(gui.GetGui(app, window))
	window.ShowAndRun()
}
