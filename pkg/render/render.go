package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Aritra640/Go_course/HTMLdemo/pkg/config"
	"github.com/Aritra640/Go_course/HTMLdemo/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	//Add data
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) error {

	//create a template cache
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var tc map[string]*template.Template
	if app.UseCache { //if template is same
		tc = app.TemplateCache
	} else { //template is changes
		tc, _ = CreateTemplateCache()
	}
	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(errors.New("error : template not found in cache"))
	}

	buf := new(bytes.Buffer) //checking for any unexpected error in the map
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		//log.Fatal(fmt.Sprintf("Some unexpected error happened in template cache , check cache\n Terminating ...\n %s", err.Error()))
		log.Fatalf("some error happened in template cache,\n %s", err.Error())
	}

	//render template

	_, err = buf.WriteTo(w)
	if err != nil {
		// return errors.New(fmt.Sprintf("Error occured in writing data(template) to io.Writer (http.ResponseWriter) , check buffer\n %s", err.Error()))
		return fmt.Errorf("error occured in writing data(template) to io.Writer(http.ResponseWriter) , check buffer\n %s", err.Error())
	}
	// parsedTemplate, err := template.ParseFiles("/media/aritra/new_volume/Go_course/HTMLdemo/templates/"+tmpl, "/media/aritra/new_volume/Go_course/HTMLdemo/templates/base.layout.tmpl")
	// if err != nil {
	// 	return errors.New("error : could not parse .tmpl file")
	// }

	// err = parsedTemplate.Execute(w, nil)

	// if err != nil {
	// 	return errors.New(fmt.Sprintf("error could not execute parsed template , %s", err.Error()))
	// }

	return nil // executed successfully
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}
	//get all files named *.page.tmpl
	pages, err := filepath.Glob("/media/aritra/new_volume/Go_course/HTMLdemo/templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}
	//range through all files with *.page.tmpl
	for _, page := range pages {
		file_name := filepath.Base(page)
		ts, err := template.New(file_name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		//check if any layout.tmpl file exist
		matches, err := filepath.Glob("/media/aritra/new_volume/Go_course/HTMLdemo/templates/*layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(matches) > 0 {
			//check templates
			ts, err = ts.ParseGlob("/media/aritra/new_volume/Go_course/HTMLdemo/templates/*layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		//add template to cache
		cache[file_name] = ts
	}
	return cache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) error {

// 	_, inMap := tc[t]
// 	if inMap {
// 		//template found in cache
// 		log.Println("Using cached template")
// 	} else {
// 		//create template
// 		log.Println("Creating template and adding to cache")
// 		err := createTemplateCache(t)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	tmpl := tc[t]
// 	err := tmpl.Execute(w, nil)

// 	return err
// }

// func createTemplateCache(t string) error {

// 	templates := []string{
// 		fmt.Sprintf("/media/aritra/new_volume/Go_course/HTMLdemo/templates/%s", t),
// 		"/media/aritra/new_volume/Go_course/HTMLdemo/templates/base.layout.tmpl",
// 	}

// 	//create the template

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil { // error in file address
// 		return err
// 	}

// 	//add template to cache

// 	tc[t] = tmpl
// 	return nil
// }
