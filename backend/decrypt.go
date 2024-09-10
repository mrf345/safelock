package backend

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	sl "github.com/mrf345/safelock-cli/safelock"
	"github.com/mrf345/safelock-cli/slErrs"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) Decrypt(path string, password string) (id string, err error) {
	if len(a.task.id) > 0 {
		log.Fatal("starting new task while in progress, race condition")
		return
	}

	var outputPath string
	var inputFile *os.File
	id = uuid.New().String()
	ctx, cancel := context.WithCancel(a.ctx)

	if outputPath, err = OpenDirectoryDialog(ctx, runtime.OpenDialogOptions{
		Title:                "Folder to decrypt files into",
		CanCreateDirectories: true,
	}); err != nil {
		cancel()
		return
	}

	if outputPath == "" {
		cancel()
		return "", err
	}

	if inputFile, err = os.Open(path); err != nil {
		a.ShowErrMsg(fmt.Sprintf("Failure: %s", err.Error()))
		cancel()
		return
	}

	a.task.cancel = cancel
	a.task.lock = sl.New()
	a.task.lock.Quiet = true
	a.task.id = id
	a.task.kind = kindDecrypt

	a.task.lock.StatusObs.
		On(sl.StatusUpdate.Str(), a.updateStatus).
		On(sl.StatusEnd.Str(), a.resetTask)

	go func() {
		if err = a.task.lock.Decrypt(ctx, inputFile, outputPath, password); err == nil {
			a.ShowInfoMsg("All set, and decrypted!")
		} else if _, invalid := errors.Unwrap(err).(*slErrs.ErrFailedToAuthenticate); invalid {
			a.ShowErrMsg("Failure: invalid password or corrupted .sla file")
		} else if !errors.Is(err, context.DeadlineExceeded) {
			a.ShowErrMsg(fmt.Sprintf("Failure: %s", err.Error()))
		}

		inputFile.Close()
		WindowSetTitle(a.ctx, Name)
	}()

	return
}
