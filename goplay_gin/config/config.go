package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PROJECT_PORT string
	PROJECT_HOST string

	POSTGRES_DB       string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_PORT     string
	POSTGRES_HOST     string
	PGSSLMODE         string

	DEFAULTOFFSET string
	DEFAULTLIMIT  string

	DOCKER_POSTGRES_CONTAINER_NAME string
}

func Load() (Config, error) {
	cfg := Config{}

	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	cfg.PROJECT_PORT = cast.ToString(getOrReturnDefaultValue("PROJECT_PORT", "3000"))
	cfg.PROJECT_HOST = cast.ToString(getOrReturnDefaultValue("PROJECT_HOST", "localhost"))

	cfg.POSTGRES_DB = cast.ToString(getOrReturnDefaultValue("POSTGRES_DB", "postgres"))
	cfg.POSTGRES_USER = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	cfg.POSTGRES_PASSWORD = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "secret"))
	cfg.POSTGRES_PORT = cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT", "5432"))
	cfg.POSTGRES_HOST = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	cfg.PGSSLMODE = cast.ToString(getOrReturnDefaultValue("PGSSLMODE", "disable"))

	cfg.DEFAULTLIMIT = cast.ToString(getOrReturnDefaultValue("DEFAULTLIMIT", "10"))
	cfg.DEFAULTOFFSET = cast.ToString(getOrReturnDefaultValue("DEFAULTOFFSET", "0"))

	cfg.DOCKER_POSTGRES_CONTAINER_NAME = cast.ToString(getOrReturnDefaultValue("DOCKER_POSTGRES_CONTAINER_NAME", "postgres_database"))

	return cfg, nil
}

// handle .env var if not exists return default val
func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
