package msgs

import tea "charm.land/bubbletea/v2"

type ViewMsg string

func ChangeView(v string) tea.Cmd {
	return func() tea.Msg {
		return ViewMsg(v)
	}
}
