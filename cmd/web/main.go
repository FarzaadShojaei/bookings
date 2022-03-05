package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tsawler/bookings/internal/config"
	"github.com/tsawler/bookings/internal/handlers"
	"github.com/tsawler/bookings/internal/models"
	"github.com/tsawler/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

//Home is Home PAge Handler

/*
//add values adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot Divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided bt %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot Divide Zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
*/
//main is the main application function
func main() {

	//what am I Going to Put in The Session
	gob.Register(models.Reservation{})
	//change this to true when in production

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot Create Template")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	//http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting Application on Post : %s", portNumber))
	//	_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
