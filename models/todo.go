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

type Config struct {
	Server struct {
		Port int    `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		DbName   string `yaml:"dbname", envconfig:"DB_NAME"`
		Username string `yaml:"dbuser", envconfig:"DB_USERNAME"`
		Password string `yaml:"password", envconfig:"DB_PASSWORD"`
	} `yaml:"database"`
}

func ProcessToJson(body []byte) (*Todo, error) {
	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		return &Todo{}, err
	}
	return &todo, nil
}
