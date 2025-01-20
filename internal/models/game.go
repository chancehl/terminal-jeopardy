package models

type JeopardyGame struct {
	Seed   string          `json:"seed"`
	Rounds []JeopardyRound `json:"rounds"`
}
