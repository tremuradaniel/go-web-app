package main

import (
	"go-web-app/pkg/config"
	"go-web-app/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.Home))

	return mux
}
