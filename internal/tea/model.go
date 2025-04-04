package internal_tea

import (
	"github.com/chancehl/terminal-jeopardy/internal/db"
	"github.com/chancehl/terminal-jeopardy/internal/models"
	"github.com/chancehl/terminal-jeopardy/internal/services"
)

type Model struct {
	game    models.JeopardyGame
	cursorX int
	cursorY int
}

var gameService services.GameService

func InitializeModel() Model {
	sqlc, err := db.GetDbConnection()
	if err != nil {
		panic(err)
	}

	dbc := db.NewDbClient(sqlc)
	questions, err := dbc.GetAllQuestions()
	if err != nil {
		panic(err)
	}

	gameService = *services.NewGameService(questions)
	game := gameService.CreateNewGame()

	return Model{
		game:    game,
		cursorX: 0,
		cursorY: 0,
	}
}
