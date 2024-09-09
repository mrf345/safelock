package backend

import (
	"fmt"

	sl "github.com/mrf345/safelock-cli/safelock"
)

func (a *App) updateStatus(status string, percent float64) {
	a.task.status = status
	a.task.percent = percent

	if percent > 0.0 {
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
	a.task = Task{}
}

func (a *App) offTaskHandlers() {
	if a.task.lock != nil {
		a.task.lock.StatusObs.
			Off(sl.StatusUpdate.Str(), a.updateStatus).
			Off(sl.StatusEnd.Str(), a.resetTask)
	}
}