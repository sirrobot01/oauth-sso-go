package routes

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/handlers"
	"github.com/sirrobot01/oauth-sso/config"
)

func UserRoutes(routers *common.Routers, cfg *config.Config) {

	routers.AllowAny.Handle(common.GetPath("user:login"), handlers.LoginHandler(cfg))
	routers.AllowAny.Handle(common.GetPath("user:register"), handlers.RegisterHandler(cfg))
}
