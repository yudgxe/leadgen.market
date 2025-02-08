package api

import (
	"github/yudgxe/leadgen.market/service/api/hash"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter() (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(chimiddleware.Recoverer)

	apiRouter := chi.NewRouter()

	hash.Bind(apiRouter)

	r.Mount("/hash", apiRouter)
	return r, nil
}
