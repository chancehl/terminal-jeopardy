package services

import (
	"fmt"

	"github.com/chancehl/terminal-jeopardy/internal/models"
)

type GameService interface{}

type gameService struct {
	game  *models.JeopardyGame
	state GameState
}

type GameState struct {
	answered map[int]bool
	round    models.JeopardyRound
	active   bool
}

// Instantiates a new game service
func NewGameService(game *models.JeopardyGame) *gameService {
	return &gameService{state: GameState{}, game: game}
}

// Starts the game and assigns default values to state
func (s *gameService) StartGame() {
	// assign default state
	s.state = GameState{
		active:   true,
		answered: make(map[int]bool),
		round:    s.game.Rounds[0],
	}

	// print round
	s.state.round.PrintRoundCategories()
}

// Gets a reference to the current round
func (s *gameService) GetCurrentRound() *models.JeopardyRound {
	return &s.state.round
}

// Marks a question as answered
func (s *gameService) AnswerQuestion(q models.JeopardyQuestion) error {
	if s.state.answered[q.Id] {
		return fmt.Errorf("question %d already answered", q.Id)
	}
	if !s.isQuestionInGame(q) {
		return fmt.Errorf("question %d does not exist in game %s", q.Id, s.game.Seed)
	}
	// TODO: validate answer
	s.state.answered[q.Id] = true
	return nil
}

// Returns true if the question exists in the current game, false otherwise
func (s *gameService) isQuestionInGame(q models.JeopardyQuestion) bool {
	for _, round := range s.game.Rounds {
		for _, category := range round.Categories {
			for _, question := range category.Questions {
				if question.Id == q.Id {
					return true
				}
			}
		}
	}
	return false
}
