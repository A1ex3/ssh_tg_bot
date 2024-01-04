package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	TgBotDebug    bool
	UserList      []int64 `json:"userList"`
	TgBotApiToken string  `json:"tgBotApiToken"`
}

func (config *Config) Unmarshal(path_to_config string) error {
	file, err := os.Open(path_to_config)

	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)

	json.Unmarshal(byteValue, &config)
	return nil
}
