package tui

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func Start() {
	program := tea.NewProgram(initialModel())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error initializing Bubble Tea, %v", err)
		os.Exit(1)
	}
}
