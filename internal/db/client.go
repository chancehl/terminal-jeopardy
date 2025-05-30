package db

import (
	"database/sql"
	"fmt"

	"github.com/chancehl/terminal-jeopardy/internal/models"
	"github.com/chancehl/terminal-jeopardy/internal/parser"
)

type DbClient struct {
	db *sql.DB
}

// Instantiates a new questions database client
func NewDbClient(db *sql.DB) *DbClient {
	return &DbClient{db}
}

// Seeds the database
func (c *DbClient) SeedDatabase() error {
	allQuestions, err := parser.ParseQuestionsJson()
	if err != nil {
		return err
	}

	err = c.CreateQuestions(allQuestions)
	if err != nil {
		return err
	}

	return nil
}

// Inserts multiple questions into the database
func (c *DbClient) CreateQuestions(questions []models.JeopardyQuestion) error {
	statement, err := c.db.Prepare(`INSERT INTO questions (id, game_id, category, round, prompt, answer, monetary_value) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return fmt.Errorf("could not prepare statement: %v", err)
	}
	defer statement.Close()

	for _, question := range questions {
		if _, err := statement.Exec(question.Id, question.GameId, question.Category, question.Round, question.Prompt, question.Answer, question.Value); err != nil {
			return err
		}
	}

	return nil
}

// Gets a question by ID
func (c *DbClient) GetQuestionById(id int) (*models.JeopardyQuestion, error) {
	var question models.JeopardyQuestion

	query := `SELECT id, game_id, category, round, prompt, answer, monetary_value FROM questions WHERE id = $1`

	if err := c.db.QueryRow(query, id).Scan(&question.Id, &question.GameId, &question.Category, &question.Round, &question.Prompt, &question.Answer, &question.Value); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &question, nil
}

// Gets all questions in database
func (c *DbClient) GetAllQuestions() ([]models.JeopardyQuestion, error) {
	var questions []models.JeopardyQuestion

	rows, err := c.db.Query("SELECT id, game_id, category, round, prompt, answer, monetary_value FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question models.JeopardyQuestion

		if err := rows.Scan(&question.Id, &question.GameId, &question.Category, &question.Round, &question.Prompt, &question.Answer, &question.Value); err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}
