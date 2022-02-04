package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id        int64  `json:"id"`
	Context   string `json:"context"`
	Owner     string `json:"owner"`
	Completed string `json:"completed"`
}

func ProcessToJson(body []byte) (*Todo, error) {
	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		return &Todo{}, err
	}
	return &todo, nil
}
