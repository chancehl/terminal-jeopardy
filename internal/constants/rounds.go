package constants

var Rounds = struct {
	Jeopardy             string
	DoubleJeopardy       string
	FinalJeopardy        string
	CategoriesPerRound   int
	QuestionsPerCategory int
}{
	Jeopardy:             "Jeopardy",
	DoubleJeopardy:       "DoubleJeopardy",
	FinalJeopardy:        "FinalJeopardy",
	CategoriesPerRound:   6,
	QuestionsPerCategory: 5,
}
