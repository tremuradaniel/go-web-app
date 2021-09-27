package main

import (
	"fmt"
	"go-web-app/pkg/config"
	"go-web-app/pkg/handlers"
	"go-web-app/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8082"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot craete template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ = means that if there is an error, pay no attention to it
	_ = http.ListenAndServe(portNumber, nil)
}
