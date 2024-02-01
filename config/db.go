package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqliteDB(dsn string) (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

func InitMysqlDB(dsn string) (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

func InitPostgresDB(dsn string) (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

func InitDB(driver string, dsn string) (db *gorm.DB) {
	fmt.Println("Connecting to database...")
	switch driver {
	case "sqlite":
		return InitSqliteDB(dsn)
	case "mysql":
		return InitMysqlDB(dsn)
	case "postgres":
		return InitPostgresDB(dsn)
	default:
		return InitSqliteDB(dsn)
	}
}
