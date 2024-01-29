package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirrobot01/oauth-sso/api/handlers"
	"github.com/sirrobot01/oauth-sso/config"
)

func UserRoutes(r chi.Router, cfg *config.Config) {
	r.Route("/user", func(r chi.Router) {
		r.Handle("/login", handlers.LoginHandler(cfg))
		r.Handle("/register", handlers.RegisterHandler(cfg))
	})
}
