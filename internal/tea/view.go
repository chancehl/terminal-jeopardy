package internal_tea

import "fmt"

func (m Model) View() string {
	s := "Model state:\n\n"

	s += fmt.Sprintf("seed: %s\n", m.game.Seed)
	s += fmt.Sprintf("cursorX: %d\n", m.cursorX)
	s += fmt.Sprintf("cursorY: %d\n", m.cursorY)

	s += "\nPress r to regenerate game or q to quit.\n"

	// Send the UI for rendering
	return s
}
