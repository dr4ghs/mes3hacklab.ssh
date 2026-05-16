// Package splashscreen
package splashscreen

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/log/v2"

	"github.com/dr4hgs/mes3hacklab.ssh/assets"
	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/home/content/navigator"
)

var splashScreenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

type Model struct {
	width  int
	height int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			nav, err := navigator.New()
			if err != nil {
				log.Error(err.Error())
				return m, nil
			}

			return nav, nil
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	return tea.NewView(
		splashScreenStyle.
			Width(m.width).Height(m.height).
			Align(lipgloss.Center, lipgloss.Center).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Center,
					assets.Banner(),
					// Enter site
					lipgloss.JoinHorizontal(
						lipgloss.Center,
						splashScreenStyle.Width(20).Render("enter"),
						splashScreenStyle.Width(20).
							AlignHorizontal(lipgloss.Right).
							Render("Start navigating"),
					),
					// Insert mode
					lipgloss.JoinHorizontal(
						lipgloss.Center,
						splashScreenStyle.Width(20).Render("i"),
						splashScreenStyle.Width(20).
							AlignHorizontal(lipgloss.Right).
							Render("Enter insert mode"),
					),
					// Help
					lipgloss.JoinHorizontal(
						lipgloss.Center,
						splashScreenStyle.Width(20).Render("?"),
						splashScreenStyle.Width(20).
							AlignHorizontal(lipgloss.Right).
							Render("Show help panel"),
					),
				),
			),
	)
}
