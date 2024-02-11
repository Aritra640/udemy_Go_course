package handlers

import (
	"net/http"

	"github.com/Aritra640/Go_course/HTMLdemo/pkg/config"
	"github.com/Aritra640/Go_course/HTMLdemo/pkg/models"
	"github.com/Aritra640/Go_course/HTMLdemo/pkg/render"
)

// Repo is the Repository used by the handler
var Repo *Repository

// TemplateData holds data sent from handler to template
// type TemplateData struct {
// 	StringMap map[string]string
// 	IntMap    map[string]int
// 	FloatMap  map[string]float32
// 	Data      map[string]interface{} // Data is a map of string that points tp 'any' other data
// 	CSRFToken string
// 	Flash     string
// 	Warning   string
// 	Error     string
// }

type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//handler for home page

	//store ip address in sessions

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//handler for about page
	stringmap := make(map[string]string)
	stringmap["test"] = "Hello!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringmap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringmap,
	})
}

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	//Writes everything to w
// 	parsedTemplate, err := template.ParseFiles("/media/aritra/new_volume/Go_course/HTMLdemo/templates/" + tmpl)
// 	if err != nil {
// 		log.Println("Could not get template", err.Error())
// 		return
// 	}
// 	err = parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		log.Println("Error occured - counld not parse HTML template ,", err.Error())
// 		return
// 	}
// }
