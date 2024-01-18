package config

import "gorm.io/gorm"

type Config struct {
	Host  string
	Port  string
	ENV   string
	Debug bool
	DB    *gorm.DB
}

func New(db *gorm.DB) *Config {
	return &Config{
		Host:  "localhost",
		Port:  "8100",
		ENV:   "dev",
		Debug: true,
		DB:    db,
	}
}
