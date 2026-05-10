// Package home
package home

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/ssh"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/home/footer"
	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/home/header"
	pm "github.com/dr4hgs/mes3hacklab.ssh/internal/tui/home/promptmode"
)

type Model struct {
	session ssh.Session
	width   int
	height  int

	mode    pm.PromptMode
	header  tea.Model
	content tea.Model
	footer  tea.Model
}

func New(s ssh.Session) tea.Model {
	mode := pm.NormalMode

	return Model{
		session: s,
		mode:    mode,
		header:  header.New(mode, s.User()),
		footer:  footer.New(mode),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.header.Init(),
		m.footer.Init(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch m.mode {
		case pm.NormalMode:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "i":
				m.mode = pm.InsertMode
				cmds = append(cmds, pm.SwitchPromptMode(m.mode))
			}
		case pm.InsertMode:
			switch msg.String() {
			case "esc":
				m.mode = pm.NormalMode
				cmds = append(cmds, pm.SwitchPromptMode(m.mode))
			}
		}
	}

	m.header, cmd = m.header.Update(msg)
	cmds = append(cmds, cmd)

	m.footer, cmd = m.footer.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() tea.View {
	var v tea.View

	cmp := lipgloss.NewCompositor(
		lipgloss.NewLayer(
			lipgloss.NewStyle().
				Width(m.width).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("10")).
				Bold(true).
				Padding(0, 1).
				Render(m.header.View().Content),
		),
		lipgloss.NewLayer(m.footer.View().Content).Y(m.height-1),
	)

	v.SetContent(cmp.Render())

	return v
}
