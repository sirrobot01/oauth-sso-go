package routes

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/handlers"
	"github.com/sirrobot01/oauth-sso/config"
)

func OauthRoutes(routers *common.Routers, cfg *config.Config) {
	routers.AllowAuth.Handle(common.GetPath("oauth:authorize"), handlers.AuthorizeHandler(cfg))
}
