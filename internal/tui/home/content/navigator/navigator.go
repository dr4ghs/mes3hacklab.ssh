// Package navigator
package navigator

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/log/v2"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/vterm"
)

type Model struct {
	interactive bool
	node        vterm.Node
	selection   int
}

func New() (Model, error) {
	node, err := vterm.Init()
	if err != nil {
		return Model{}, err
	}

	return Model{
		interactive: true,
		node:        node,
		selection:   0,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.selection = (m.selection + 1) % len(m.node.Children)
		case "k", "up":
			m.selection--
			if m.selection < 0 {
				m.selection = len(m.node.Children) - 1
			}
		case "l", "enter":
			if node, err := vterm.Cd(m.node, m.node.Children[m.selection].Name); err != nil {
				log.Error(err.Error())
			} else {
				m.node = node
			}
		case "h", "backspace":
			if node, err := vterm.Cd(m.node, ".."); err != nil {
				log.Error(err.Error())
			} else {
				m.node = node
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() tea.View {
	var v tea.View

	v.SetContent(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.RenderElems()...,
		),
	)

	return v
}

func (m Model) RenderElems() (res []string) {
	ls := vterm.Ls(m.node, true)

	for i, e := range ls {
		if m.selection == i {
			res = append(
				res,
				lipgloss.NewStyle().
					Foreground(lipgloss.Color("10")).
					Bold(true).
					Render(e),
			)
		} else {
			res = append(res, e)
		}
	}

	return
}
