package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetGui(a fyne.App, w fyne.Window) fyne.CanvasObject {
	act := Action{App: a, Window: w}

	dark := widget.NewToolbarAction(theme.VisibilityOffIcon(), act.switchLight)
	dark.ToolbarObject().Hide()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.VisibilityIcon(), act.switchDark),
		dark,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), act.showAboutDialog),
	)
	act.ToolBar = toolbar

	img := canvas.NewImageFromResource(resourceLogoPng)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(logoSize)

	footer := container.NewVBox(
		widget.NewButton("Encrypt folder", act.selectFolderAndEncrypt),
		widget.NewButton("Decrypt file", act.selectFileAndDecrypt),
		widget.NewLabel(""),
	)
	progBar := widget.NewProgressBar()
	progLabel := widget.NewLabel("")
	progBar.Min = 0.0
	progBar.Max = 100.0
	progLabel.Alignment = fyne.TextAlignCenter
	busyFooter := container.NewVBox(
		progLabel,
		progBar,
		widget.NewButton("Cancel", act.cancelProcess),
	)
	body := container.NewVBox(
		toolbar,
		img,
		footer,
		busyFooter,
	)
	busyFooter.Hidden = true
	act.Footer = footer
	act.BusyFooter = busyFooter
	act.ProgBar = progBar
	act.ProgLabel = progLabel

	return body
}
