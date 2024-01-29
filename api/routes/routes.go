package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/middlewares"
	"github.com/sirrobot01/oauth-sso/config"
	"net/http"
)

func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RealIP)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	authRouter := r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware(cfg))
	})

	// Add Static File Server
	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("welcome"))
		if err != nil {
			return
		}
	})

	authRouter.Get(common.GetPath(":welcome"), func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("This is the special route with custom middleware"))
		if err != nil {
			return
		}
		return
	})

	// Register routes
	UserRoutes(r, authRouter, cfg)

	return r
}
