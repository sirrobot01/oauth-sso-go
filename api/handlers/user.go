package handlers

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/models"
	"github.com/sirrobot01/oauth-sso/api/schemas"
	"github.com/sirrobot01/oauth-sso/api/services"
	"github.com/sirrobot01/oauth-sso/config"
	"net/http"
	"strconv"
	"time"
)

//func LoginHandler(w http.ResponseWriter, r *http.Request) {
//	var data LoginInSchema
//	err := common.ReaderToStruct(r, &data)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	common.RenderJSON(w, data, http.StatusOK)
//}

func RegisterHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()

			if err != nil {
				return
			}
			username := r.Form.Get("username")
			password := r.Form.Get("password")
			passwordConfirm := r.Form.Get("confirm_password")
			data := schemas.RegisterInSchema{
				Username:        username,
				Password:        password,
				ConfirmPassword: passwordConfirm,
			}
			err = data.Validate()
			if err != nil {
				err := common.RenderTemplate(w, "register.html", err)
				if err != nil {
					// re-render the template
					_ = common.RenderTemplate(w, "register.html", nil)
					return
				}
			}
			service := services.New(cfg)
			_, err = service.RegisterUser(&data)
			if err != nil {
				err := common.RenderTemplate(w, "register.html", err)
				if err != nil {
					// re-render the template
					_ = common.RenderTemplate(w, "register.html", nil)
					return
				}
				return
			}
			//common.RenderJSON(w, user, http.StatusOK)
			http.Redirect(w, r, common.GetPath("user:login"), http.StatusFound)
			return
		}
		err := common.RenderTemplate(w, "register.html", nil)
		if err != nil {
			return
		}
	}
}

func LoginHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Render the HTML template
		if r.Method == http.MethodPost {
			err := r.ParseForm()

			if err != nil {
				return
			}
			username := r.Form.Get("username")
			password := r.Form.Get("password")
			data := schemas.LoginInSchema{
				Username: username,
				Password: password,
			}
			err = data.Validate()
			if err != nil {
				err := common.RenderTemplate(w, "login.html", err)
				if err != nil {
					// re-render the template
					_ = common.RenderTemplate(w, "login.html", nil)
					return
				}
			}
			service := services.New(cfg)
			var user *models.User
			user, err = service.LoginUser(&data)
			if err != nil {
				err := common.RenderTemplate(w, "login.html", err)
				if err != nil {
					// re-render the template
					_ = common.RenderTemplate(w, "login.html", nil)
					return
				}
				return
			}
			//Set the cookie
			token, _ := common.GenerateToken(strconv.Itoa(int(user.ID)), user.Username, user.IsAdmin)
			expires := time.Now().Add(time.Minute * 100)
			common.SetCookie(w, "trk", token, expires)
			http.Redirect(w, r, "/welcome", http.StatusFound)
			return
		}
		err := common.RenderTemplate(w, "login.html", nil)
		if err != nil {
			return
		}

		return

	}
}
