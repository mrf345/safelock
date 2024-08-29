package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	slErrs "github.com/mrf345/safelock-cli/errors"
	sl "github.com/mrf345/safelock-cli/safelock"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) Encrypt(paths []string, password string) (id string, err error) {
	if len(a.task.id) > 0 {
		log.Fatal("starting new task while in progress, race condition")
		return
	}

	var outputPath string
	id = uuid.New().String()
	ctx, cancel := context.WithCancel(a.ctx)

	if outputPath, err = SaveFileDialog(ctx, runtime.SaveDialogOptions{
		DefaultFilename:      "encrypted.sla",
		Title:                "Save encrypted .sla file to",
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
	a.task.kind = kindEncrypt

	a.task.lock.StatusObs.
		On(sl.EventStatusUpdate, a.updateStatus).
		On(sl.EventStatusEnd, a.resetTask)

	go func() {
		if err = a.task.lock.Encrypt(ctx, paths, outputPath, password); err == nil {
			a.ShowInfoMsg("All set, and encrypted!")
		} else if _, cancelled := err.(*slErrs.ErrContextExpired); !cancelled {
			a.ShowErrMsg(fmt.Sprintf("Failure: %s", err.Error()))
		}
	}()

	return
}
