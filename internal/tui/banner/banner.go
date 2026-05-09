package banner

import (
	_ "embed"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/ssh"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/msgs"
)

//go:embed banner.txt
var banner string

type model struct {
	term          string
	profile       string
	viewportWidth int
	progress      int
	width         int
	height        int
	bg            string
	txtStyle      lipgloss.Style
}

func New(pty ssh.Pty) tea.Model {
	vpWidth := lipgloss.Width(banner)

	m := model{
		term:          pty.Term,
		viewportWidth: vpWidth,
		progress:      vpWidth,
		width:         pty.Window.Width,
		height:        pty.Window.Height,
		bg:            "ligt",
		txtStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("10")),
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var c []tea.Cmd

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
			c = append(c, msgs.ChangeView("landing"))
		}
	case msgs.TickMsg:
		m.progress--
		if m.progress < -100 {
			c = append(c, msgs.ChangeView("landing"))
		}

		c = append(c, msgs.Tick(time.Millisecond))
	}

	return m, tea.Batch(c...)
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
				Height(m.height).
				Width(m.progress).
				Align(lipgloss.Center, lipgloss.Center).
				Render(""),
		).
			X((m.width-m.viewportWidth)/2+(m.viewportWidth-m.progress)).
			Z(10),
	)
	v.SetContent(cmp.Render())

	return v
}
