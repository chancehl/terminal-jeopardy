package main

import (
	"fmt"
	"os"

	internal_tea "github.com/chancehl/terminal-jeopardy/internal/tea"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()
	}

	p := tea.NewProgram(internal_tea.InitializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("an error happened: %v", err)
		os.Exit(1)
	}
}
