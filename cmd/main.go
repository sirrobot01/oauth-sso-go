package main

import (
	"fmt"
	"github.com/sirrobot01/oauth-sso/api/routes"
	"github.com/sirrobot01/oauth-sso/config"
	"net/http"
)

func main() {

	cfg := config.InitConfig()
	router := routes.NewRouter(cfg)

	fmt.Println("Running server on " + "http://" + cfg.Host + ":" + cfg.Port)
	err := http.ListenAndServe(cfg.Host+":"+cfg.Port, router)
	if err != nil {
		return
	}
}
