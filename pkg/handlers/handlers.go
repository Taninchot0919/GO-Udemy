package handlers

import (
	"net/http"

	"github.com/taninchot0919/learning-go/pkg/config"
	"github.com/taninchot0919/learning-go/pkg/render"
)

var Repo *Respository

type Respository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		App: a,
	}
}

func NewHandlers(r *Respository)  {
	Repo = r
	
}

func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
