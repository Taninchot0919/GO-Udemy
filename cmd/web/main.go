package main

import (
	"log"
	"net/http"

	"github.com/taninchot0919/learning-go/pkg/config"
	"github.com/taninchot0919/learning-go/pkg/handlers"
	"github.com/taninchot0919/learning-go/pkg/render"
)

const portNumber = ":9000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(portNumber, nil)
}
