package gui

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	myErrs "github.com/mrf345/safelock-cli/errors"
)

type Action struct {
	App            fyne.App
	Window         fyne.Window
	ToolBar        *widget.Toolbar
	Footer         *fyne.Container
	BusyFooter     *fyne.Container
	ProgBar        *widget.ProgressBar
	ProgLabel      *widget.Label
	Cancel         context.CancelFunc
	encryptChannel chan EncryptItem
}

type EncryptItem struct {
	DirPath  string
	FilePath string
	Pwd      string
}

func (a *Action) showAboutDialog() {
	ghUrl, _ := url.Parse(Github)
	label := widget.NewLabel(About)
	url := widget.NewHyperlink(Github+"\n", ghUrl)
	label.Alignment = fyne.TextAlignCenter
	url.Alignment = fyne.TextAlignCenter
	body := container.NewVBox(label, url)
	dialog.NewCustom("About", "Close", body, a.Window).Show()
}

func (a *Action) cancelProcess() {
	if a.Cancel == nil {
		a.handleProgressErr(errors.New("failed to cancel, no context found"))
		return
	}

	a.Cancel()
	a.Cancel = nil
}

func (a *Action) handleProgressErr(err error) {
	if _, cancelled := (err).(*myErrs.ErrContextExpired); !cancelled {
		dialog.ShowError(err, a.Window)
	}

	a.BusyFooter.Hide()
	a.Footer.Show()
}

func (a *Action) validatePassword(s string) (err error) {
	if 8 > len(strings.TrimSpace(s)) {
		err = errors.New("must be at leas 8 characters (no white spaces)")
	}
	return
}

func (a *Action) handleSuccessDialog(msg string) dialog.Dialog {
	info := dialog.NewInformation("Success", msg, a.Window)
	info.SetOnClosed(func() {
		a.ProgBar.SetValue(0)
		a.ProgLabel.SetText("")
		a.BusyFooter.Hide()
		a.Footer.Show()
	})
	info.Show()
	return info
}

func (a *Action) updateProgressBar(status string, percent float64) {
	a.ProgLabel.SetText(status)
	a.ProgBar.SetValue(percent)
}

func (a *Action) switchDark() {
	a.App.Settings().SetTheme(GetCustomTheme(true))
	a.ToolBar.Items[0].ToolbarObject().Hide()
	a.ToolBar.Items[1].ToolbarObject().Show()
}

func (a *Action) switchLight() {
	a.App.Settings().SetTheme(GetCustomTheme(false))
	a.ToolBar.Items[1].ToolbarObject().Hide()
	a.ToolBar.Items[0].ToolbarObject().Show()
}
