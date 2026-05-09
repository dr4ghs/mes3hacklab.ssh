package landing

import (
	tea "charm.land/bubbletea/v2"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui/msgs"
)

type model struct{}

func New() tea.Model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			return m, msgs.ChangeView("banner")
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	var v tea.View

	v.SetContent("Hello, world!")

	return v
}
