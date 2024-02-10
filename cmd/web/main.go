package main

import (
	"log"
	"net/http"
	"time"

	handlers "github.com/Aritra640/Go_course/HTMLdemo/pkg"
	"github.com/Aritra640/Go_course/HTMLdemo/pkg/config"
	"github.com/Aritra640/Go_course/HTMLdemo/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = ":8000"

// func Home(w http.ResponseWriter, r *http.Request) {
// 	//handler for home page
// 	renderTemplate(w, "home.page.tmpl")
// }

// func About(w http.ResponseWriter, r *http.Request) {
// 	//handler for about page
// 	renderTemplate(w, "about.page.tmpl")
// }

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

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		// log.Fatal(fmt.Sprintf("error cannot create template cache \nterminating...\n %s", err.Error()))
		log.Fatalf("error cannot create template cache\n terminating....\n %s", err.Error())
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // if the browser is closed the data will persist -> true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// session.Cookie.Secure = false // in production this should be true

	// var app config.AppConfig
	app.TemplateCache = tc
	app.UseCache = false

	app.InProduction = false

	session.Cookie.Secure = app.InProduction
	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	log.Println("Starting Port ", PORT)

	// err = http.ListenAndServe(PORT, nil)

	// if err != nil {
	// 	log.Fatal(errors.New("error : failed to start server"))
	// }

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error : cannot handle routes , terminating...\n%s", err.Error())
	}

}
