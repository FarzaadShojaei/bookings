package handlers

import (
	"net/http"

	"github.com/tsawler/bookings/pkg/config"
	"github.com/tsawler/bookings/pkg/models"
	"github.com/tsawler/bookings/pkg/render"
)

//Template Data holds data sent from handlers to templates

//the repository used by a handler
var Repo *Repository

//is the repository type
type Repository struct {
	App *config.AppConfig
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

//creates a New Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//About is about Page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello,again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

/*
//	fmt.Fprintf(w, "This is The Home Page")
//	sum := addValues(2, 2)
	//	_, _ = fmt.Fprintf(w, fmt.Sprintf("This Is The About Page and 2+2 is : %d", sum))
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing Template", err)
	}
}
*/
