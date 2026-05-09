package tui

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/ssh"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/banner"
	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/landing"
	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/msgs"
)

func Handler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()

	m := model{
		banner:  banner.New(pty),
		landing: landing.New(),
	}

	return m, []tea.ProgramOption{}
}

type model struct {
	banner  tea.Model
	landing tea.Model
	current *tea.Model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		msgs.ChangeView("banner"),
		tea.RequestBackgroundColor,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var c []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			c = append(c, tea.Quit)
		}
	case msgs.ViewMsg:
		switch msg {
		case "banner":
			m.current = &m.banner
			c = append(c, msgs.Tick(time.Millisecond))
		case "landing":
			m.current = &m.landing
		}
	}

	if m.current != nil {
		model, cmd := (*m.current).Update(msg)
		*m.current = model

		c = append(c, cmd)
	}

	return m, tea.Batch(c...)
}

func (m model) View() tea.View {
	var v tea.View

	if m.current != nil {
		v = (*m.current).View()
	}
	v.AltScreen = true

	return v
}
