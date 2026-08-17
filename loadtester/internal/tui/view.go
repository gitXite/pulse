package tui

import (
	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	s := "placeholder"

	// for i, choice := range m.choices {
	// 	cursor := " "
		
	// 	if m.cursor == i {
	// 		cursor = ">"
	// 	}
	// }

	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}