package generators

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"

	"github.com/chancehl/terminal-jeopardy/pkg/constants"
	"github.com/chancehl/terminal-jeopardy/pkg/models"
)

// Generates a seed for the provided rounds of jeopardy
func generateSeed(rounds []models.JeopardyRound) string {
	questions := []models.JeopardyQuestion{}

	for _, value := range rounds {
		for _, category := range value.Categories {
			questions = append(questions, category.Questions...)
		}
	}

	ids := make([]string, len(questions))

	for index, question := range questions {
		ids[index] = strconv.Itoa(question.Id)
	}

	seed := base64.StdEncoding.EncodeToString([]byte(strings.Join(ids, ",")))

	return seed
}

// Groups questions by round
func groupByRound(questions []models.JeopardyQuestion) map[string][]models.JeopardyQuestion {
	grouped := make(map[string][]models.JeopardyQuestion)

	for _, question := range questions {
		grouped[question.Round] = append(grouped[question.Round], question)
	}

	return grouped
}

// Picks a random question category
func pickRandomQuestionCategory(questions []models.JeopardyQuestion) []models.JeopardyQuestion {
	categoryQuestions := []models.JeopardyQuestion{}

	randomIndex := rand.Intn(len(questions))
	randomQuestion := questions[randomIndex]

	for _, question := range questions {
		if question.GameId == randomQuestion.GameId && question.Category == randomQuestion.Category {
			categoryQuestions = append(categoryQuestions, question)
		}
	}

	return categoryQuestions
}

// Picks a random question
func pickRandomQuestion(questions []models.JeopardyQuestion) models.JeopardyQuestion {
	randomIndex := rand.Intn(len(questions))
	randomQuestion := questions[randomIndex]

	return randomQuestion
}

// Picks random questions
func pickRandomQuestions(validQuestions []models.JeopardyQuestion, isFinalJeopardy bool) []models.JeopardyQuestion {
	questions := []models.JeopardyQuestion{}

	if isFinalJeopardy {
		questions = append(questions, pickRandomQuestion(validQuestions))
	} else {
		categories := 0

		for categories < constants.Rounds.CategoriesPerRound {
			questions = append(questions, pickRandomQuestionCategory(validQuestions)...)

			categories++
		}
	}

	return questions
}

// Generates categories from questions
func generateCategoriesFromQuestions(questions []models.JeopardyQuestion) []models.JeopardyCategory {
	categoryMap := make(map[string][]models.JeopardyQuestion)

	for _, q := range questions {
		categoryMap[q.Category] = append(categoryMap[q.Category], q)
	}

	categories := make([]models.JeopardyCategory, 0, len(categoryMap))

	for name, qs := range categoryMap {
		categories = append(categories, models.JeopardyCategory{Name: name, Questions: qs})
	}

	return categories
}

// Generates a random round of Jeopardy
func generateRounds(questions []models.JeopardyQuestion) []models.JeopardyRound {
	grouped := groupByRound(questions)

	jeopardyRoundQuestions := pickRandomQuestions(grouped[constants.Rounds.Jeopardy], false)
	doubleJeopardyRoundQuestions := pickRandomQuestions(grouped[constants.Rounds.DoubleJeopardy], false)
	finalJeopardyRoundQuestions := pickRandomQuestions(grouped[constants.Rounds.FinalJeopardy], true)

	jeopardyRound := models.JeopardyRound{
		Name:       constants.Rounds.Jeopardy,
		Categories: generateCategoriesFromQuestions(jeopardyRoundQuestions),
	}

	doubleJeopardyRound := models.JeopardyRound{
		Name:       constants.Rounds.DoubleJeopardy,
		Categories: generateCategoriesFromQuestions(doubleJeopardyRoundQuestions),
	}

	finalJeopardyRound := models.JeopardyRound{
		Name:       constants.Rounds.FinalJeopardy,
		Categories: generateCategoriesFromQuestions(finalJeopardyRoundQuestions),
	}

	return []models.JeopardyRound{jeopardyRound, doubleJeopardyRound, finalJeopardyRound}
}

// Generates a random Jeopardy game with random categories, questions
func GenerateGame(questions []models.JeopardyQuestion) models.JeopardyGame {
	rounds := generateRounds(questions)
	seed := generateSeed(rounds)

	return models.JeopardyGame{Rounds: rounds, Seed: seed}
}
