package config

import (
	"github.com/joho/godotenv"
	storage "github.com/sirrobot01/oauth-sso/storage/db"
	"os"
)

type Config struct {
	Host  string
	Port  string
	ENV   string
	Debug bool
	DB    *storage.DB
}

func GetEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = fallback
	}
	return value
}

func InitConfig() *Config {
	env := GetEnv("ENV", "dev")
	err := godotenv.Load("./envs/." + env + ".env")
	if err != nil {
		panic("Error loading .env file")
	}
	dsn := GetEnv("DB_DSN", "./bin/data.db")

	driver := GetEnv("DB_DRIVER", "sqlite") // sqlite, mysql, postgres
	db := InitDB(driver, dsn)

	database := &storage.DB{
		Tx: db,
	}
	return &Config{
		Host:  GetEnv("HOST", "localhost"),
		Port:  GetEnv("PORT", "8100"),
		ENV:   env,
		Debug: GetEnv("DEBUG", "true") == "true",
		DB:    database,
	}
}
