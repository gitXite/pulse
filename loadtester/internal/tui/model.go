package tui

import (
	"time"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	padding  = 2
	maxWidth = 80
)

var style = lipgloss.NewStyle().Foreground(lipgloss.Color("#5A2132")).Render

// placeholder model
type model struct {
	choices  []string
	cursor   int
	progress progress.Model
}

type tickMsg time.Time

func initialModel() model {
	return model{
		choices: []string{
			"Test",
			"Hello",
		},
	}
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
