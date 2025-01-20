package models

import "fmt"

type JeopardyCategory struct {
	Name      string             `json:"name"`
	Questions []JeopardyQuestion `json:"questions"`
}

func (c *JeopardyCategory) Print() {
	fmt.Println(c.Name)
}
