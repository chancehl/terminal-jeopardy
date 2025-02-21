package internal_tea

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q":
			return m, tea.Quit
		case "r":
			m.game = Service.CreateNewGame()
		case "up":
			m.cursorY++
		case "down":
			m.cursorY--
		case "left":
			m.cursorX--
		case "right":
			m.cursorX++
		}
	}
	return m, nil
}
