package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chancehl/terminal-jeopardy/pkg/models"
)

func ParseQuestionsJson() ([]models.JeopardyQuestion, error) {
	content, err := os.ReadFile("./questions.json")
	if err != nil {
		return nil, fmt.Errorf("could not read questions.json file: %v", err)
	}

	var questions []models.JeopardyQuestion
	json.Unmarshal(content, &questions)

	return questions, nil
}
