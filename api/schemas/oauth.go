package schemas

import "github.com/sirrobot01/oauth-sso/api/common"

var ResponseType = []string{"code", "token"}

type AuthorizeInSchema struct {
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

func (authorizeSchema *AuthorizeInSchema) Validate() error {
	if authorizeSchema.ClientID == "" {
		return &common.ValidationError{
			Message: "client_id is required",
		}
	}

	if authorizeSchema.ResponseType != "" {
		if !common.Contains(ResponseType, authorizeSchema.ResponseType) {
			return &common.ValidationError{
				Message: "invalid response_type",
			}
		}
	}

	return nil
}

type AuthorizeOutSchema struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
