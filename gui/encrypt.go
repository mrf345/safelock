package gui

import (
	"context"
	"errors"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/mrf345/safelock-cli/safelock"
)

func (a *Action) selectFolderAndEncrypt() {
	a.encryptChannel = make(chan EncryptItem, 1)
	dialog.ShowFolderOpen(a.handleFolderOpen, a.Window)
}

func (a *Action) handleFolderOpen(lu fyne.ListableURI, err error) {
	if lu == nil {
		return
	} else if err != nil {
		dialog.ShowError(err, a.Window)
		return
	}

	dialog.ShowFileSave(a.handleFileSaveAndPassword, a.Window)
	a.encryptChannel <- EncryptItem{DirPath: filepath.FromSlash(lu.Path())}
}

func (a *Action) handleFileSaveAndPassword(uc fyne.URIWriteCloser, err error) {
	if uc == nil {
		return
	} else if err != nil {
		dialog.ShowError(err, a.Window)
		return
	}

	pwd := widget.NewPasswordEntry()
	pwd2 := widget.NewPasswordEntry()
	pwd.Validator = a.validatePassword
	pwd2.Validator = func(s string) (err error) {
		if s != pwd.Text {
			err = errors.New("must match password")
		}
		return
	}
	items := []*widget.FormItem{
		{Text: "Password", Widget: pwd},
		{Text: "Confirm", Widget: pwd2},
	}

	item := <-a.encryptChannel
	item.FilePath = filepath.FromSlash(uc.URI().Path())

	form := dialog.NewForm("Enter Password", "Encrypt", "Cancel", items, func(complete bool) {
		if !complete {
			return
		}

		item.Pwd = pwd.Text
		go a.encryptItem(item)
	}, a.Window)

	form.Show()
	form.Resize(formSize)
	close(a.encryptChannel)
}

func (a *Action) encryptItem(item EncryptItem) {
	var ctx context.Context

	a.Footer.Hide()
	a.BusyFooter.Show()
	a.ProgLabel.SetText(item.FilePath)

	ctx, a.Cancel = context.WithCancel(context.TODO())
	lock := safelock.New()
	lock.Quiet = true

	lock.StatusObs.On(safelock.EventStatusUpdate, a.updateProgressBar)

	if err := lock.Encrypt(ctx, item.DirPath, item.FilePath, item.Pwd); err != nil {
		a.handleProgressErr(err)
		lock.StatusObs.Off(safelock.EventStatusUpdate, a.updateProgressBar)
		return
	}

	a.handleSuccessDialog("Encrypted folder successfully")
}
