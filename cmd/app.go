package main

import (
	"fmt"
	"os"

	internal_tea "github.com/chancehl/terminal-jeopardy/internal/tea"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(internal_tea.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("an error happened: %v", err)
		os.Exit(1)
	}
}
