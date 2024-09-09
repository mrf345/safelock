package backend

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/mrf345/safelock-cli/safelock"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	Version         = "1.0.0"
	Name            = "Safelock"
	statusUpdateKey = "status_update"
	statusEndKey    = "status_end"
	openedSlaKey    = "opened_sla_file"
	kindEncrypt     = "encrypt"
	kindDecrypt     = "decrypt"
)

var (
	MessageDialog       = runtime.MessageDialog
	SaveFileDialog      = runtime.SaveFileDialog
	OpenDirectoryDialog = runtime.OpenDirectoryDialog
	EventsEmit          = runtime.EventsEmit
)

type Task struct {
	id      string
	status  string
	percent float64
	kind    string
	lock    *safelock.Safelock
	cancel  context.CancelFunc
}

type App struct {
	ctx  context.Context
	task Task
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	runtime.WindowCenter(ctx)

	isSlaFileOpened := len(os.Args) > 1 && strings.HasSuffix(os.Args[1], ".sla")

	if isSlaFileOpened {
		go func() {
			time.Sleep(time.Second / 3)
			EventsEmit(ctx, openedSlaKey, os.Args[1])
			runtime.WindowShow(ctx)
		}()
	}
}

func (a *App) GetVersion() string {
	return Version
}

func (a *App) ShowErrMsg(msg string) {
	_, _ = MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "😞 Failure",
		Message: msg,
	})
}

func (a *App) ShowInfoMsg(msg string) {
	_, _ = MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "🎉 Success",
		Message: msg,
	})
}

func (a *App) Cancel() {
	if len(a.task.id) > 0 {
		a.task.cancel()
	}
}
