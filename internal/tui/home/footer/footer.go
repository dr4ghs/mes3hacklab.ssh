// Package footer
package footer

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	pm "github.com/dr4hgs/mes3hacklab.ssh/internal/tui/home/mode"
)

const _NormalModeHintText = "Press ? for help"

var hintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

type Model struct {
	width int

	mode pm.PromptMode
}

func New(mode pm.PromptMode) Model {
	return Model{
		mode: mode,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case pm.PromptMode:
		m.mode = msg
	}

	return m, nil
}

func (m Model) View() tea.View {
	hint := hintStyle.AlignHorizontal(lipgloss.Right).Render(_NormalModeHintText)
	if m.mode != pm.NormalMode {
		hint = ""
	}

	cmp := lipgloss.NewCompositor(
		lipgloss.NewLayer(fmt.Sprintf("-- %s --", m.mode.String())).X(0),
		lipgloss.NewLayer(hint).X(m.width-lipgloss.Width(hint)),
	)

	return tea.NewView(cmp.Render())
}

func (m Model) SetWidth(width int) Model {
	m.width = width

	return m
}
