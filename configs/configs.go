package configs

import (
	"os"

	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/joho/godotenv"
)

type Config struct {
}

func New(path string) *Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Error("Failed to load config", err.Error())
	}

	log.Info("Config loaded from path", path)
	return &Config{}
}

func (c *Config) Get(key string) string {
	return os.Getenv(key)
}
