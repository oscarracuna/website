package handlers

import (
	"net/http"

	"github.com/oscarracuna/website/pkg/config"
	"github.com/oscarracuna/website/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Functions that start with Capital Letter are global
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", nil)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", nil)
}
