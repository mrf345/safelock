package gui

import (
	"context"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/mrf345/safelock-cli/safelock"
)

func (a *Action) selectFileAndDecrypt() {
	a.encryptChannel = make(chan EncryptItem, 1)
	dialog.ShowFileOpen(a.handleFileOpen, a.Window)
}

func (a *Action) handleFileOpen(uc fyne.URIReadCloser, err error) {
	if uc == nil {
		return
	} else if err != nil {
		dialog.ShowError(err, a.Window)
		return
	}

	dialog.ShowFolderOpen(a.handleDecryptFileSaveAndPassword, a.Window)
	a.encryptChannel <- EncryptItem{FilePath: filepath.FromSlash(uc.URI().Path())}
}

func (a *Action) handleDecryptFileSaveAndPassword(lu fyne.ListableURI, err error) {
	if lu == nil {
		return
	} else if err != nil {
		dialog.ShowError(err, a.Window)
		return
	}

	pwd := widget.NewPasswordEntry()
	pwd.Validator = a.validatePassword
	items := []*widget.FormItem{{Text: "Password", Widget: pwd}}
	item := <-a.encryptChannel
	item.DirPath = filepath.FromSlash(lu.Path())

	form := dialog.NewForm("Enter Password", "Decrypt", "Cancel", items, func(complete bool) {
		if !complete {
			return
		}

		item.Pwd = pwd.Text
		go a.decryptItem(item)
	}, a.Window)

	form.Show()
	form.Resize(formSize)
	close(a.encryptChannel)
}

func (a *Action) decryptItem(item EncryptItem) {
	var ctx context.Context

	a.Footer.Hide()
	a.BusyFooter.Show()
	a.ProgLabel.SetText(item.FilePath)

	ctx, a.Cancel = context.WithCancel(context.TODO())
	lock := safelock.New()
	lock.Quiet = true

	lock.StatusObs.On(safelock.EventStatusUpdate, a.updateProgressBar)

	if err := lock.Decrypt(ctx, item.FilePath, item.DirPath, item.Pwd); err != nil {
		a.handleProgressErr(err)
		lock.StatusObs.Off(safelock.EventStatusUpdate, a.updateProgressBar)
		return
	}

	a.handleSuccessDialog("Decrypted file successfully")
}
