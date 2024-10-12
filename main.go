package main

import (
	"log"

	"github.com/chancehl/terminal-jeopardy/pkg/db"
	"github.com/chancehl/terminal-jeopardy/pkg/generators"
	"github.com/chancehl/terminal-jeopardy/pkg/services"
)

func main() {
	dbConnection, err := db.GetDbConnection()
	if err != nil {
		log.Fatalf("could not create database connection: %v", err)
	}
	defer dbConnection.Close()

	// if err := db.SeedDatabase(); err != nil {
	// 	log.Fatalf("could not seed database: %v", err)
	// }

	dbClient := db.NewDbClient(dbConnection)

	questions, err := dbClient.GetAllQuestions()
	if err != nil {
		log.Fatalf("could not retrieve questions: %v", err)
	}

	// instantiate game data and service
	game := generators.GenerateGame(*questions)
	gameService := services.NewGameService(game)

	// start the game
	gameService.StartGame()
}
