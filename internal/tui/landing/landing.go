// Package landing
package landing

import (
	_ "embed"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/ssh"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/home"
	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/msgs"
)

//go:embed banner.txt
var banner string

func Banner() string {
	return banner
}

type model struct {
	session       ssh.Session
	term          string
	profile       string
	viewportWidth int
	progress      int
	width         int
	height        int
	bg            string
	txtStyle      lipgloss.Style
	helpStyle     lipgloss.Style
}

func New(s ssh.Session) tea.Model {
	pty, _, _ := s.Pty()
	vpWidth := lipgloss.Width(banner)

	m := model{
		session:       s,
		term:          pty.Term,
		viewportWidth: vpWidth,
		progress:      vpWidth,
		width:         pty.Window.Width,
		height:        pty.Window.Height,
		bg:            "ligt",
		txtStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("10")),
		helpStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.ColorProfileMsg:
		m.profile = msg.String()
	case tea.BackgroundColorMsg:
		if msg.IsDark() {
			m.bg = "dark"
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m.Next()
		}
	case msgs.TickMsg:
		m.progress--
		if m.progress < -100 {
			return m.Next()
		}

		cmds = append(cmds, msgs.Tick(time.Millisecond))
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() tea.View {
	var v tea.View

	cmp := lipgloss.NewCompositor(
		lipgloss.NewLayer(
			m.txtStyle.Width(m.width).
				Height(m.height).
				Align(lipgloss.Center, lipgloss.Center).
				Render(banner),
		).Z(0),
		lipgloss.NewLayer(
			lipgloss.NewStyle().
				Height(m.height-1).
				Width(m.progress).
				Align(lipgloss.Center, lipgloss.Center).
				Render(""),
		).
			X((m.width-m.viewportWidth)/2+(m.viewportWidth-m.progress)).
			Z(10),
		lipgloss.NewLayer(
			m.helpStyle.Align(lipgloss.Center, lipgloss.Center).
				Width(m.width).
				Render("Press 'enter' to skip"),
		).Y(m.height-1).Z(0),
	)
	v.SetContent(cmp.Render())

	return v
}

func (m model) Next() (tea.Model, tea.Cmd) {
	next := home.New(m.session)

	return next, tea.Batch(
		next.Init(),
		func() tea.Msg {
			return tea.WindowSizeMsg{
				Width:  m.width,
				Height: m.height,
			}
		},
	)
}
