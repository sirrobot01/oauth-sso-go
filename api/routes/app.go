package routes

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/handlers"
	"github.com/sirrobot01/oauth-sso/config"
)

func AppRoutes(routers *common.Routers, cfg *config.Config) {

	routers.AllowAuth.Handle(common.GetPath("app:new"), handlers.CreateAppHandler(cfg))
	routers.AllowAuth.Handle(common.GetPath("app:list"), handlers.ListAppsHandler(cfg))
}
