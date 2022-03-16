package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tsawler/bookings/internal/config"
	"github.com/tsawler/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	testApp.UseCache = true
	app = &testApp

	os.Exit(m.Run())
}

// myWriter will satisfy the response writer interface
// by implementing its methods, and we can use this
// to test rendering templates
type myWriter struct{}

// Header is used to satisfy the http.ResponseWriter interface
func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

// Write is used to satisfy the http.ResponseWriter interface
// Note that we must return the length of the bytes written
func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

// WriteHeader is used to satisfy the http.ResponseWriter interface
func (tw *myWriter) WriteHeader(s int) {}
