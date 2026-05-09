package msgs

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type TickMsg time.Time

func Tick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
