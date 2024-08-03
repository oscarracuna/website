package main

// TODO Figure out HTMX
// TODO Simplify landing page
// TODO Find better color palette

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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("App starting on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
