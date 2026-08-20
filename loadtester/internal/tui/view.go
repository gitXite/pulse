package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	s := "Pulse"
	pad := strings.Repeat(" ", padding)

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("\n%s %s \n", cursor, choice)
	}

	return tea.NewView("\n" +
		s + "\n" +
		pad + m.progress.View() + "\n\n" +
		pad + style("Press q to quit."))
}
