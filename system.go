package main

import (
	"log"
	"os"
)

// ConfigData contains config data. Api keys, paths
type ConfigData struct {
	botAPIKey     string
	dbPath        string
}

func loadConfigFromEnv() ConfigData {

	botAPIKey, isSet := os.LookupEnv("TELEGRAM_BOT_API_KEY_VOICE")
	if !isSet {
		log.Panic("No bot api key found. Please set TELEGRAM_BOT_API_KEY_VOICE env")
	}

	dbPath, isSet := os.LookupEnv("TELEGRAM_BOT_DB_PATH")
	if !isSet {
		log.Panic("No bot api key found. Please set TELEGRAM_BOT_DB_PATH env")
	}

	log.Println(botAPIKey + " " + dbPath)

	return ConfigData{botAPIKey, dbPath}
}
