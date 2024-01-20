package user

import (
	"github.com/sirrobot01/oauth-sso/config"
	"github.com/sirrobot01/oauth-sso/internal/common"
	"net/http"
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
			data := RegisterInSchema{
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
			data := LoginInSchema{
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
			return
		}
		err := common.RenderTemplate(w, "login.html", nil)
		if err != nil {
			return
		}
	}
}
