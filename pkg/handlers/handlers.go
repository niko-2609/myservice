package handlers

import (
	"net/http"

	"github.com/myservice/pkg/config"
	"github.com/myservice/pkg/models"
	"github.com/myservice/pkg/render"
)

// Repo is a variable that will have many configuration objects
var Repo *Repository

//Repository is the type of object to hold all the configuration objects inside
/** Ex:
     Repo = {
		{app1: {<AppConfig for this app1 here>}},
		{app2: {<AppConfig for this app2 here>}},
		{app3: {<AppConfig for this app3 here>}},
	} */

type Repository struct {
	App *config.AppConfig
}

// Get the config from main and return an Repository type object containing the received config
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Get a Repository type object from main and set it to Repo variable to use in this package
func NewHandlers(r *Repository) {
	Repo = r
}

/**
m is an pointer receiver to the type Repository.
These receivers are added to functions to allow those functions have access to the object directly.
Also all the methods using these receivers, will have instant access to any change made to this object by any other method.
Thus all transactions are managed by Go and its safe to use and also best practice.
*/
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.html.tmpl", &models.TemplateData{})

}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login-page.html.tmpl", &models.TemplateData{})

}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, world"
	render.RenderTemplate(w, "about-page.html.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
