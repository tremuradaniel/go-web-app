package main

import (
	"fmt"
	"go-web-app/internal/config"
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing: test passed
	default:
		t.Error(fmt.Sprintf("type is not *Chi.Mux, type is %T", v))
	}
}
