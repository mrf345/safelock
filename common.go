package main

import (
	"context"
	"time"

	"github.com/mrf345/safelock-cli/safelock"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	Version         = "1.0.0"
	Name            = "Safelock"
	statusUpdateKey = "status_update"
	statusEndKey    = "status_end"
	kindEncrypt     = "encrypt"
	kindDecrypt     = "decrypt"
)

var (
	MessageDialog       = runtime.MessageDialog
	SaveFileDialog      = runtime.SaveFileDialog
	OpenDirectoryDialog = runtime.OpenDirectoryDialog
	EventsEmit          = runtime.EventsEmit
)

type Pacer struct {
	duration time.Duration
	ready    bool
}

func (p *Pacer) pace() {
	time.Sleep(p.duration)
	p.ready = true
}

func (p *Pacer) Ready() bool {
	if p.ready {
		p.ready = false
		go p.pace()
		return true
	}

	return p.ready
}

type Task struct {
	id      string
	status  string
	percent float64
	kind    string
	lock    *safelock.Safelock
	cancel  context.CancelFunc
}

type App struct {
	ctx   context.Context
	task  Task
	pacer *Pacer
}
