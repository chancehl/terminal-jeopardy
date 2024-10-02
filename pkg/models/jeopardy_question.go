package models

type JeopardyQuestion struct {
	Prompt   string `json:"prompt"`
	Category string `json:"category"`
	Round    string `json:"round"`
	Value    int    `json:"value"`
	Answer   string `json:"answer"`
	Id       int    `json:"id"`
	GameId   int    `json:"gameId"`
}
