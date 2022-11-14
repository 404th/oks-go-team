package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PORT string
	HOST string

	PGDATABASE string
	PGUSER     string
	PGPASSWORD string
	PGPORT     string
	PGHOST     string
	PGSSLMODE  string

	DEFAULTOFFSET string
	DEFAULTLIMIT  string
}

func Load() (Config, error) {
	cfg := Config{}

	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	cfg.PORT = cast.ToString(getOrReturnDefaultValue("PORT", "3000"))
	cfg.HOST = cast.ToString(getOrReturnDefaultValue("HOST", "localhost"))

	cfg.PGDATABASE = cast.ToString(getOrReturnDefaultValue("PGDATABASE", "postgres"))
	cfg.PGUSER = cast.ToString(getOrReturnDefaultValue("PGUSER", "postgres"))
	cfg.PGPASSWORD = cast.ToString(getOrReturnDefaultValue("PGPASSWORD", "secret"))
	cfg.PGPORT = cast.ToString(getOrReturnDefaultValue("PGPORT", "5432"))
	cfg.PGHOST = cast.ToString(getOrReturnDefaultValue("PGHOST", "localhost"))
	cfg.PGSSLMODE = cast.ToString(getOrReturnDefaultValue("PGSSLMODE", "disable"))

	cfg.DEFAULTLIMIT = cast.ToString(getOrReturnDefaultValue("DEFAULTLIMIT", "10"))
	cfg.DEFAULTOFFSET = cast.ToString(getOrReturnDefaultValue("DEFAULTOFFSET", "0"))

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
