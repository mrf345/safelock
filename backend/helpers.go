package backend

import (
	"fmt"
	"strings"

	sl "github.com/mrf345/safelock-cli/safelock"
)

func (a *App) updateStatus(status string, percent float64) {
	a.task.status = status
	a.task.percent = percent

	if percent > 0.0 {
		WindowSetTitle(
			a.ctx, fmt.Sprintf(
				"%sing (%.2f%%)",
				strings.Title(a.task.kind.Str()), //nolint:all
				percent,
			),
		)
		EventsEmit(
			a.ctx,
			statusUpdateKey,
			status,
			fmt.Sprintf("%.2f", percent),
		)
	}
}

func (a *App) resetTask() {
	a.offTaskHandlers()
	EventsEmit(a.ctx, statusEndKey)
	WindowSetTitle(a.ctx, Name)
	a.task = Task{}
}

func (a App) offTaskHandlers() {
	if a.task.lock != nil {
		a.task.lock.StatusObs.
			Off(sl.StatusUpdate.Str(), a.updateStatus).
			Off(sl.StatusEnd.Str(), a.resetTask)
	}
}
