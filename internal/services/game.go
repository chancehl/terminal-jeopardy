package services

import (
	"fmt"

	"github.com/chancehl/terminal-jeopardy/internal/db"
	"github.com/chancehl/terminal-jeopardy/internal/generators"
	"github.com/chancehl/terminal-jeopardy/internal/models"
)

type gameService struct {
	dbc   *db.DbClient
	game  models.JeopardyGame
	state GameState
}

type GameState struct {
	answered []int
	round    models.JeopardyRound
}

// Instantiates a new game service
func NewGameService(dbc *db.DbClient) *gameService {
	return &gameService{state: GameState{}, dbc: dbc}
}

// Generates a game
func (s *gameService) CreateGame() error {
	questions, err := s.dbc.GetAllQuestions()
	if err != nil {
		return fmt.Errorf("could not retrieve questions: %v", err)
	}

	s.game = generators.GenerateGame(questions)
	return nil
}

// Starts the game and assigns default values to state
func (s *gameService) StartGame() {
	s.state = GameState{
		answered: []int{},
		round:    s.game.Rounds[0],
	}
}

// Gets a reference to the  current round
func (s *gameService) GetCurrentRound() *models.JeopardyRound {
	return &s.state.round
}

// Marks a question as answered
func (s *gameService) AnswerQuestion(q models.JeopardyQuestion) error {
	if !s.isQuestionInGame(q) {
		return fmt.Errorf("question %d does not exist in game %s", q.Id, s.game.Seed)
	}
	// TODO: validate answer
	s.state.answered = append(s.state.answered, q.Id)
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
