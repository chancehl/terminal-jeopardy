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

	// err = db.SeedDatabase(db.RoundCounts{Jeopardy: 500, DoubleJeopardy: 500, FinalJeopardy: 50})
	// if err != nil {
	// 	log.Fatalf("could not seed database: %v", err)
	// } else {
	// 	fmt.Println("seeded database")
	// }

	dbClient := db.NewDbClient(dbConnection)

	questions, err := dbClient.GetQuestions()
	if err != nil {
		log.Fatalf("could not retrieve questions: %v", err)
	}

	// instantiate game data and service
	game := generators.GenerateGame(*questions)
	gameService := services.NewGameService(game)

	// start the game
	gameService.StartGame()
}
