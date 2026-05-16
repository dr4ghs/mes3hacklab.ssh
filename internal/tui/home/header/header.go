// Package header
package header

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	pm "github.com/dr4hgs/mes3hacklab.ssh/internal/tui/home/mode"
)

var promptLineStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10")).
	AlignHorizontal(lipgloss.Left)

type Model struct {
	mode pm.PromptMode

	identity  string
	home      string
	path      string
	textInput textinput.Model
}

func New(mode pm.PromptMode, identity string) Model {
	ti := textinput.New()
	ti.Prompt = promptLineStyle.Render("# ")
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.SetWidth(50)

	return Model{
		mode:      mode,
		home:      "/home/mhl",
		path:      "",
		identity:  fmt.Sprintf("%s@mes3hacklab", identity),
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case pm.PromptMode:
		m.mode = msg
		m.textInput.SetVirtualCursor(m.mode == pm.InsertMode)
		m.textInput.Placeholder = ""
		if m.mode == pm.InsertMode {
			m.textInput.Placeholder = "help"
		}

		if m.mode == pm.NormalMode {
			m.textInput.SetValue("")
		}
	case tea.KeyMsg:
		if m.mode == pm.InsertMode {
			switch msg.String() {
			case "tab":
				if m.textInput.Value() == "" {
					m.textInput.SetValue(m.textInput.Placeholder)
				}
			case "enter":
				line := m.textInput.Value()
				m.textInput.Reset()

				if line == "exit" {
					return m, tea.Quit
				}
			case "ctrl+d":
				return m, tea.Quit
			}

			m.textInput, cmd = m.textInput.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() tea.View {
	var v tea.View

	sb := strings.Builder{}
	sb.WriteString(promptLineStyle.Render(fmt.Sprintf("%s:%s%s ", m.identity, m.home, m.path)))
	sb.WriteString(promptLineStyle.Render(m.textInput.View()))

	v.SetContent(sb.String())

	return v
}
