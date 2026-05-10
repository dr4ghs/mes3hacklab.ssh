// Package mode
package mode

import tea "charm.land/bubbletea/v2"

type PromptMode string

func (m PromptMode) String() string {
	switch m {
	case NormalMode:
		return "NORMAL"
	case InsertMode:
		return "INSERT"
	}

	return "UNKNOWN"
}

var (
	NormalMode PromptMode = "n"
	InsertMode PromptMode = "i"
)

func SwitchPromptMode(mode PromptMode) tea.Cmd {
	return func() tea.Msg {
		return mode
	}
}
