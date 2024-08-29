package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	slErrs "github.com/mrf345/safelock-cli/errors"
	sl "github.com/mrf345/safelock-cli/safelock"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) Decrypt(path string, password string) (id string, err error) {
	if len(a.task.id) > 0 {
		log.Fatal("starting new task while in progress, race condition")
		return
	}

	var outputPath string
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

	a.task.cancel = cancel
	a.task.lock = sl.New()
	a.task.lock.Quiet = true
	a.task.id = id
	a.task.kind = kindDecrypt

	a.task.lock.StatusObs.
		On(sl.EventStatusUpdate, a.updateStatus).
		On(sl.EventStatusEnd, a.resetTask)

	go func() {
		if err = a.task.lock.Decrypt(ctx, path, outputPath, password); err == nil {
			a.ShowInfoMsg("All set, and decrypted!")
		} else if _, invalid := errors.Unwrap(err).(*slErrs.ErrFailedToAuthenticate); invalid {
			a.ShowErrMsg("Failure: invalid password or corrupted .sla file")
		} else if _, cancelled := err.(*slErrs.ErrContextExpired); !cancelled {
			a.ShowErrMsg(fmt.Sprintf("Failure: %s", err.Error()))
		}
	}()

	return
}
