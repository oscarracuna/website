package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oscarracuna/website/pkg/config"
	"github.com/oscarracuna/website/pkg/handlers"
	"github.com/oscarracuna/website/pkg/render"
)

const (
	portNumber = ":8888"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Unable to create template cache")
	}

	app.TemplateCache = tc
	// This right here determines if we're in developer mode or prod
	// false = dev mode so we can make changes and see them immediately
	// true = prod
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("App starting on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
