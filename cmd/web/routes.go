package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/dev-ayaa/track-space/pkg/config"
	"github.com/dev-ayaa/track-space/pkg/handler"

)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", handler.HomePage)

	return mux
}
