package main

import (
	"fmt"
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
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
