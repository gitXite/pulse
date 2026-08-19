package tui

import (
	tea "charm.land/bubbletea/v2"
)

// placeholder model
type model struct {
	choices []string
	cursor int
}

func initialModel() model {
	return model{
		choices: []string{""},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}