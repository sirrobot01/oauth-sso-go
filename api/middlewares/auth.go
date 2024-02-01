package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/models"
	"github.com/sirrobot01/oauth-sso/config"
	"net/http"
)

func AuthMiddleware(cfg *config.Config) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenCookie, err := common.GetCookie(r, "trk")
			if err != nil {
				AuthFailed(w, r)
				return
			}
			token, err := jwt.ParseWithClaims(tokenCookie.Value, &common.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.GetEnv("SECRET_KEY", "secret")), nil
			})
			if err != nil {
				AuthFailed(w, r)
				return
			}
			claims, ok := token.Claims.(*common.JwtClaims)
			if !ok {
				AuthFailed(w, r)
				return
			}
			//Validate
			if err := claims.Valid(); err != nil {
				AuthFailed(w, r)
				return
			}
			//if err := common.ValidateToken(claims, cfg); err != nil {
			//	AuthFailed(w, r)
			//	return
			//}
			ctx := r.Context()
			_, ok = ctx.Value("user").(*models.User)
			if !ok {
				// cached user not in context
				user, err := cfg.DB.GetUserById(claims.UserID)
				if err != nil {
					AuthFailed(w, r)
					return
				}
				ctx = context.WithValue(r.Context(), "user", user)

			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})

	}
}

func AuthFailed(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, common.GetPath("user:login"), http.StatusSeeOther)
}

func AnonymousOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := common.GetCookie(r, "trk")
		if err == nil && tokenCookie != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
