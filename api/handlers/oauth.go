package handlers

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/schemas"
	"github.com/sirrobot01/oauth-sso/api/services"
	"github.com/sirrobot01/oauth-sso/config"
	"net/http"
	"strings"
)

func AuthorizeHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ...
		queries := r.URL.Query()
		params := schemas.AuthorizeInSchema{
			ClientID:     queries.Get("client_id"),
			RedirectURI:  queries.Get("redirect_uri"),
			ResponseType: queries.Get("response_type"),
			Scope:        queries.Get("scope"),
			State:        queries.Get("state"),
		}

		err := params.Validate()
		if err != nil {
			err := common.RenderTemplate(w, "authorize.html", err)
			if err != nil {
				// re-render the template
				_ = common.RenderTemplate(w, "authorize.html", nil)
				return
			}
			return
		}

		service := services.New(cfg)
		client, err := service.GetClient(params.ClientID)
		if err != nil {
			err := common.RenderTemplate(w, "authorize.html", err)
			if err != nil {
				// re-render the template
				_ = common.RenderTemplate(w, "authorize.html", nil)
				return
			}
			return
		}
		if client == nil {
			err = &common.ValidationError{
				Message: "client not found",
			}
			_ = common.RenderTemplate(w, "authorize.html", err)
			return
		}
		clientRedirectURIs := strings.Split(client.RedirectURIs, ",")
		if params.RedirectURI != "" {
			if !common.Contains(clientRedirectURIs, params.RedirectURI) {
				err = &common.ValidationError{
					Message: "redirect uri not allowed",
				}
				_ = common.RenderTemplate(w, "authorize.html", err)
				return
			}
		} else {
			params.RedirectURI = clientRedirectURIs[0]
		}

		tmplContext := make(map[string]any)
		tmplContext["clientName"] = client.Name
		tmplContext["redirectURI"] = params.RedirectURI

		err = common.RenderTemplate(w, "authorize.html", tmplContext)
		if err != nil {
			return
		}
		if r.Method == http.MethodPost {
			//err := r.ParseForm()
			//if err != nil {
			//	return
			//}
			//params.ClientID = r.Form.Get("client_id")
			//params.RedirectURI = r.Form.Get("redirect_uri")

			app, err := service.AuthorizeClient(client, &params)
			if err != nil {
				err := common.RenderTemplate(w, "authorize.html", err)
				if err != nil {
					// re-render the template
					_ = common.RenderTemplate(w, "authorize.html", nil)
					return
				}
				return
			}
			schemaOut := schemas.AuthorizeOutSchema{
				Code:  app.AuthCode,
				State: params.State,
			}
			common.RenderJSON(w, schemaOut, http.StatusOK)
			return
		}
		return
	}

}
