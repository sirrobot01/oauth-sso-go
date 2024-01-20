package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirrobot01/oauth-sso/config"
)

func Routes(r chi.Router, cfg *config.Config) {
	r.Route("/user", func(r chi.Router) {
		r.Handle("/login", LoginHandler(cfg))
		r.Handle("/register", RegisterHandler(cfg))
	})
}
