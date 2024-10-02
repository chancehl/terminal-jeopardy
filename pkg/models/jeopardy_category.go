package models

type JeopardyCategory struct {
	Name      string             `json:"name"`
	Questions []JeopardyQuestion `json:"questions"`
}
