package configs

import (
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

func (c *configLoader) Load() *Config {
	appEnv := os.Getenv("PROJECT_ENV")
	if appEnv == "" {
		godotenv.Load(".env")
	} else {
		godotenv.Load(appEnv + ".env")
	}

	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
