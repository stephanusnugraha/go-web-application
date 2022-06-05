package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	return mux
}
