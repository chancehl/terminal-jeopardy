package services

import (
	"fmt"

	"github.com/chancehl/terminal-jeopardy/internal/generators"
	"github.com/chancehl/terminal-jeopardy/internal/models"
)

type GameService struct {
	game      models.JeopardyGame
	questions []models.JeopardyQuestion
	state     GameState
}

type GameState struct {
	answered []int
	round    models.JeopardyRound
}

// Instantiates a new game service
func NewGameService(questions []models.JeopardyQuestion) *GameService {
	return &GameService{questions: questions}
}

// Generates a game
func (s *GameService) CreateNewGame() models.JeopardyGame {
	s.game = generators.GenerateGame(s.questions)
	return s.game
}

// Starts the game and assigns default values to state
func (s *GameService) StartGame() {
	s.state = GameState{
		answered: []int{},
		round:    s.game.Rounds[0],
	}
}

// Gets a reference to the  current round
func (s *GameService) GetCurrentRound() *models.JeopardyRound {
	return &s.state.round
}

// Marks a question as answered
func (s *GameService) AnswerQuestion(q models.JeopardyQuestion) error {
	if !s.isQuestionInGame(q) {
		return fmt.Errorf("question %d does not exist in game %s", q.Id, s.game.Seed)
	}
	// TODO: validate answer
	s.state.answered = append(s.state.answered, q.Id)
	return nil
}

// Returns true if the question exists in the current game, false otherwise
func (s *GameService) isQuestionInGame(q models.JeopardyQuestion) bool {
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
