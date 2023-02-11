package config

import (
	"encoding/json"
	"log"
	"os"
)

var config *Config

type Config struct {
	Openai_API_Key     string `json:"openai_api_key"`
	Telegram_Bot_Token string `json:"telegram_bot_token"`
}

// init config
func InitConfig() {
	// Open our jsonFile
	configFile, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	log.Println("config initialized")
	log.Println("openai api key:", config.Openai_API_Key)
	log.Println("telegram bot token:", config.Telegram_Bot_Token)
}

func GetConfig() *Config {
	return config
}
