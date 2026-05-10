// Package tui
package tui

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/ssh"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/landing"
	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/msgs"
)

func Handler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	m := model{
		current: landing.New(s),
	}

	return m, []tea.ProgramOption{}
}

type model struct {
	current tea.Model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		msgs.Tick(time.Millisecond),
		tea.RequestBackgroundColor,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.current, cmd = m.current.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() tea.View {
	v := m.current.View()
	v.AltScreen = true

	return v
}
