package internal_tea

import (
	"github.com/chancehl/terminal-jeopardy/internal/constants"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q":
			return m, tea.Quit
		case "r":
			m.game = gameService.CreateNewGame()
		case "up":
			if m.cursorY < constants.Rounds.QuestionsPerCategory {
				m.cursorY++
			}
		case "down":
			if m.cursorY > 0 {
				m.cursorY--
			}
		case "left":
			if m.cursorX > 0 {
				m.cursorX--
			}
		case "right":
			if m.cursorX < constants.Rounds.CategoriesPerRound {
				m.cursorX++
			}
		}
	}
	return m, nil
}
