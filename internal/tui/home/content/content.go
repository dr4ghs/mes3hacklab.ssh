// Package content
package content

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/dr4hgs/mes3hacklab.ssh/assets"
)

type Model struct {
	value string
}

func New(val string) Model {
	return Model{
		value: val,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() tea.View {
	var v tea.View

	v.SetContent(m.value)

	return v
}

var splashScreenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

func SplashScreen() string {
	return splashScreenStyle.Render(lipgloss.JoinVertical(
		lipgloss.Center,
		assets.Banner(),
		// Enter site
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			splashScreenStyle.Width(20).Render("enter"),
			splashScreenStyle.Width(20).AlignHorizontal(lipgloss.Right).Render("Start navigating"),
		),
		// Insert mode
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			splashScreenStyle.Width(20).Render("i"),
			splashScreenStyle.Width(20).AlignHorizontal(lipgloss.Right).Render("Enter insert mode"),
		),
		// Help
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			splashScreenStyle.Width(20).Render("?"),
			splashScreenStyle.Width(20).AlignHorizontal(lipgloss.Right).Render("Show help panel"),
		),
	))
}
