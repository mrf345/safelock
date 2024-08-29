package main

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func NewApp() *App {
	return &App{
		pacer: &Pacer{
			ready:    true,
			duration: time.Second / 5,
		},
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	runtime.WindowCenter(ctx)
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
