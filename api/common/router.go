package common

import "github.com/go-chi/chi/v5"

type Routers struct {
	AllowAny   *chi.Mux
	AllowAuth  chi.Router
	AllowAdmin chi.Router
}
