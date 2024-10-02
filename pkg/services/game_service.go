package services

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/chancehl/terminal-jeopardy/pkg/constants"
	"github.com/chancehl/terminal-jeopardy/pkg/models"
)

type GameService interface{}

type gameService struct {
	game  models.JeopardyGame
	state GameState
}

type GameState struct {
	answered map[int]bool
	round    string
	active   bool
}

// Instantiates a new game service
func NewGameService(game models.JeopardyGame) *gameService {
	return &gameService{state: GameState{}, game: game}
}

// Starts the game and assigns default values to state
func (s *gameService) StartGame() {
	s.state = GameState{
		active:   true,
		answered: make(map[int]bool),
		round:    constants.Rounds.Jeopardy,
	}
}

// Gets a reference to the current round
func (s *gameService) GetCurrentRound() (*models.JeopardyRound, error) {
	for _, round := range s.game.Rounds {
		if round.Name == s.state.round {
			return &round, nil
		}
	}
	return nil, fmt.Errorf("could not locate current round")
}

// Prints the categories to the std out
func (s *gameService) PrintCategories(round models.JeopardyRound) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.DiscardEmptyColumns)

	var headers []string
	for _, category := range round.Categories {
		headers = append(headers, category.Name)
	}

	header := strings.Join(headers, "\t")

	fmt.Fprintln(writer, header)

	var values []int

	if round.Name == constants.Rounds.Jeopardy {
		values = []int{200, 400, 600, 800, 1000}
	} else if round.Name == constants.Rounds.DoubleJeopardy {
		values = []int{400, 800, 1200, 1600, 2000}
	} else {
		values = []int{-1} // TODO: implement wager
	}

	for _, value := range values {
		for i := 0; i < len(round.Categories); i++ {
			fmt.Fprintf(writer, "%d\t", value)
		}
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
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
