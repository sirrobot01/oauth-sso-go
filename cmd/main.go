package main

import (
	"fmt"
	"github.com/sirrobot01/oauth-sso/config"
	"github.com/sirrobot01/oauth-sso/internal/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func InitDB() *gorm.DB {
	fmt.Println("Connecting to database...")
	db, err := gorm.Open(sqlite.Open("./bin/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func main() {
	db := InitDB()
	cfg := config.New(db)
	router := routes.NewRouter(cfg)

	fmt.Println("Running server on " + cfg.Host + ":" + cfg.Port)
	err := http.ListenAndServe(cfg.Host+":"+cfg.Port, router)
	if err != nil {
		return
	}
}
