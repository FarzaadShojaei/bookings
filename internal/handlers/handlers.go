package handlers

import (
	"fmt"
	"net/http"

	"github.com/tsawler/bookings/internal/config"
	"github.com/tsawler/bookings/internal/models"
	"github.com/tsawler/bookings/internal/render"
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

	render.RenderTemplate(w, r,"home.page.tmpl",&models.TemplateData{})

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

	render.RenderTemplate(w,r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
//Reservation renders the make a Reservation page and displays form
func(m *Repository) Reservation(w http.ResponseWriter,r *http.Request){
	render.RenderTemplate(w,r,"make-reservation.page.tmpl",&models.TemplateData{})
}
//Generals renders The Room Page
func(m *Repository) Generals(w http.ResponseWriter,r *http.Request){
	render.RenderTemplate(w,r,"generals.page.tmpl",&models.TemplateData{})
}
//Majors renders The Room page
func(m *Repository) Majors(w http.ResponseWriter,r *http.Request){
	render.RenderTemplate(w,r,"majors.page.tmpl",&models.TemplateData{})
}

func(m *Repository) Availability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"search-availability.page.tmpl",&models.TemplateData{})
}

func(m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request){
	start:=r.Form.Get("start")
	end:=r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s",start,end)))
}

func(m *Repository) Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"contact.page.tmpl",&models.TemplateData{})
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