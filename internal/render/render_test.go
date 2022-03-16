package render

import (
	"net/http"
	"testing"

	"github.com/tsawler/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("not in session")
	}
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Fail()
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	// create a var that satisfies the interface for http.ResponseWriter
	var ww myWriter
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error rendering template")
	}

	err = RenderTemplate(&ww, r, "nonexistent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("got template that does not exist")
	}
}

// this creates a request with a session in the context
func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-page", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(r.Context(), r.Header.Get("X-Session"))

	r = r.WithContext(ctx)

	return r, nil
}
