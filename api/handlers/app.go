package handlers

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/models"
	"github.com/sirrobot01/oauth-sso/api/schemas"
	"github.com/sirrobot01/oauth-sso/api/services"
	"github.com/sirrobot01/oauth-sso/config"
	"net/http"
)

func CreateAppHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Render the HTML template
		if r.Method == http.MethodPost {
			user, ok := r.Context().Value("user").(*models.User)
			if !ok {
				// user not in context/ not logged in
				http.Redirect(w, r, common.GetPath("user:login"), http.StatusFound)
				return
			}

			err := r.ParseForm()
			if err != nil {
				return
			}
			name := r.Form.Get("name")
			redirectURI := r.Form.Get("redirect_uris")

			appRequest := schemas.AppRequest{
				Name:         name,
				RedirectUris: redirectURI,
				Scopes:       "read,write",
				UserID:       user.ID,
			}
			err = appRequest.Validate()
			if err != nil {
				err := common.RenderTemplate(w, "create_app.html", err)
				if err != nil {
					// re-render the template
					_ = common.RenderTemplate(w, "create_app.html", nil)
					return
				}
				return
			}

			service := services.New(cfg)
			_, err = service.CreateApp(&appRequest)
			if err != nil {
				err := common.RenderTemplate(w, "create_app.html", err)
				if err != nil {
					// re-render the template
					_ = common.RenderTemplate(w, "create_app.html", nil)
					return
				}
				return
			}
			http.Redirect(w, r, common.GetPath("app:list"), http.StatusFound)
			return
		}
		err := common.RenderTemplate(w, "create_app.html", nil)
		if err != nil {
			return
		}
		return
	}
}

func ListAppsHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Apps List"))
		if err != nil {
			return
		}
		return
	}
}
