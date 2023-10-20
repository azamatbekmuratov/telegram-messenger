package config

import (
	"os"

	"github.com/rs/zerolog/log"
)

var (
	BotToken = getEnv("BOT_TOKEN", "")
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Warn().Msgf("Using default value for %s", key)
	return defaultVal
}
