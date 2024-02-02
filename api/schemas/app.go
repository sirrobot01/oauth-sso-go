package schemas

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"net/url"
	"strings"
)

type AppRequest struct {
	Name         string `json:"name"`
	RedirectUris string `json:"redirect_uris"`
	Scopes       string `json:"scopes"`
	UserID       uint   `json:"user_id"`
}

func (createApp *AppRequest) Validate() error {
	if createApp.Name == "" {
		return &common.ValidationError{
			Message: "name is required",
		}
	}
	redirectUris := strings.Split(createApp.RedirectUris, ",")
	if len(redirectUris) == 0 {
		return &common.ValidationError{
			Message: "redirect_uris is required",
		}
	}

	// Validate redirect uris
	for _, uri := range redirectUris {
		println(uri)
		if _, err := url.ParseRequestURI(uri); err != nil {
			return &common.ValidationError{
				Message: "invalid redirect URI: " + uri,
			}
		}
	}

	return nil
}
