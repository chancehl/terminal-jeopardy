package models

type JeopardyRound struct {
	Name       string             `json:"name"`
	Categories []JeopardyCategory `json:"categories"`
}
