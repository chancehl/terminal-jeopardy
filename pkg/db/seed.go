package db

import (
	"log"

	"github.com/chancehl/terminal-jeopardy/pkg/constants"
	"github.com/chancehl/terminal-jeopardy/pkg/models"
	"github.com/chancehl/terminal-jeopardy/pkg/parser"
	"github.com/chancehl/terminal-jeopardy/pkg/utils"
)

type RoundCounts struct {
	Jeopardy       int
	DoubleJeopardy int
	FinalJeopardy  int
}

// Seeds the database
func SeedDatabase(counts RoundCounts) error {
	allQuestions, err := parser.ParseQuestionsJson()
	if err != nil {
		return err
	}

	db, err := GetDbConnection()
	if err != nil {
		log.Fatalf("could not create database connection: %v", err)
	}
	defer db.Close()

	questions := pickQuestionsByRoundCount(allQuestions, counts)

	if err := NewDbClient(db).CreateQuestions(questions); err != nil {
		return err
	}
	return nil
}

// Selects questions based on round by count
func pickQuestionsByRoundCount(questions []models.JeopardyQuestion, counts RoundCounts) []models.JeopardyQuestion {
	var selected []models.JeopardyQuestion

	questionMap := map[string][]models.JeopardyQuestion{
		constants.Rounds.Jeopardy:       make([]models.JeopardyQuestion, 0),
		constants.Rounds.DoubleJeopardy: make([]models.JeopardyQuestion, 0),
		constants.Rounds.FinalJeopardy:  make([]models.JeopardyQuestion, 0),
	}

	utils.ShuffleSlice(questions)

	for _, question := range questions {
		questionMap[question.Round] = append(questionMap[question.Round], question)
	}

	jeopardyQuestionCount := min(len(questionMap[constants.Rounds.Jeopardy]), counts.Jeopardy)
	doubleJeopardyQuestionCount := min(len(questionMap[constants.Rounds.DoubleJeopardy]), counts.DoubleJeopardy)
	finalJeopardyQuestionCount := min(len(questionMap[constants.Rounds.Jeopardy]), counts.FinalJeopardy)

	selected = append(selected, questionMap[constants.Rounds.Jeopardy][0:jeopardyQuestionCount]...)
	selected = append(selected, questionMap[constants.Rounds.DoubleJeopardy][0:doubleJeopardyQuestionCount]...)
	selected = append(selected, questionMap[constants.Rounds.FinalJeopardy][0:finalJeopardyQuestionCount]...)

	return selected
}
