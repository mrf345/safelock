package backend

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	sl "github.com/mrf345/safelock-cli/safelock"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) Encrypt(paths []string, password string) (id string, err error) {
	if len(a.task.id) > 0 {
		log.Fatal("starting new task while in progress, race condition")
		return
	}

	var outputPath string
	var outputFile *os.File
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

	fileFlags := os.O_RDWR | os.O_CREATE | os.O_TRUNC

	if outputFile, err = os.OpenFile(outputPath, fileFlags, 0755); err != nil {
		a.ShowErrMsg(fmt.Sprintf("Failure: %s", err.Error()))
		cancel()
		return
	}

	a.task.cancel = cancel
	a.task.lock = sl.New()
	a.task.lock.Quiet = true
	a.task.id = id
	a.task.kind = kindEncrypt

	a.task.lock.StatusObs.Subscribe(a.handleStatusUpdate)

	go func() {
		if err = a.task.lock.Encrypt(ctx, paths, outputFile, password); err == nil {
			a.ShowInfoMsg("All set, and encrypted!")
			outputFile.Close()
		} else if !errors.Is(err, context.DeadlineExceeded) {
			a.ShowErrMsg(fmt.Sprintf("Failure: %s", err.Error()))
			outputFile.Close()
			_ = os.Remove(outputFile.Name())
		}

		WindowSetTitle(a.ctx, Name)
	}()

	return
}
