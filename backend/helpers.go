package backend

import (
	"fmt"
	"strings"

	sl "github.com/mrf345/safelock-cli/safelock"
)

func (a *App) handleStatusUpdate(si sl.StatusItem) {
	switch si.Event {
	case sl.StatusEnd:
		a.resetTask()
	case sl.StatusUpdate:
		a.updateStatus(si)
	}
}

func (a *App) updateStatus(si sl.StatusItem) {
	kind := a.task.kind.Str()

	if kind != "" && si.Percent > 0.0 {
		a.task.status = si.Msg
		a.task.percent = si.Percent

		EventsEmit(
			a.ctx,
			statusUpdateKey,
			si.Msg,
			fmt.Sprintf("%.2f", si.Percent),
		)
		WindowSetTitle(
			a.ctx, fmt.Sprintf(
				"%sing (%.2f%%)",
				strings.Title(kind), //nolint:all
				si.Percent,
			),
		)
	}
}

func (a *App) resetTask() {
	WindowSetTitle(a.ctx, Name)
	EventsEmit(a.ctx, statusEndKey)
	a.task.lock.StatusObs.Unsubscribe()
	a.task = Task{}
}
