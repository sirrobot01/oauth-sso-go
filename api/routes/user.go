package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/handlers"
	"github.com/sirrobot01/oauth-sso/config"
)

func UserRoutes(r chi.Router, authRouter chi.Router, cfg *config.Config) {
	_ = authRouter

	r.Handle(common.GetPath("user:login"), handlers.LoginHandler(cfg))
	r.Handle(common.GetPath("user:register"), handlers.RegisterHandler(cfg))
}
