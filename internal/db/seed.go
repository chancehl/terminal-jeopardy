package db

import (
	"log"

	"github.com/chancehl/terminal-jeopardy/internal/parser"
)

// Seeds the database
func SeedDatabase() error {
	allQuestions, err := parser.ParseQuestionsJson()
	if err != nil {
		return err
	}

	db, err := GetDbConnection()
	if err != nil {
		log.Fatalf("could not create database connection: %v", err)
	}
	defer db.Close()

	if err := NewDbClient(db).CreateQuestions(allQuestions); err != nil {
		return err
	}
	return nil
}
