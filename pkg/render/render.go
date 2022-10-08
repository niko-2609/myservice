package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/myservice/pkg/config"
	"github.com/myservice/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// get the cache map by calling createTemplate function
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// check if the received tmpl name is present in the cache that is initialized
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Required template not found in the cache")
	}

	// create a buffer
	// This is kind of a replacement for the actual http Response Writer
	buf := new(bytes.Buffer)

	// Instead of directly writing the template on response writer,
	// its written in this buffer.
	// The second argument is any data that we want to pass on to the template. This would yeild a result of templates and template data and write to the buffer.
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// if everything goes well and the template execution is successfull
	// then we write the bytes present in the buffer to the http response writer
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// This function is called to create a template cache at the time the application
// is initialized
func CreateTemplateCache() (map[string]*template.Template, error) {

	// initialize cache
	myCache := map[string]*template.Template{}

	// parse through all the files in a directory
	// returns slice of full path(address of the file on the system) of the file
	pages, err := filepath.Glob("./templates/*.html.tmpl")
	if err != nil {
		return myCache, err
	}

	// loop through all the pages rin the slice received from the above code
	for _, page := range pages {

		// filepath.Base gives the end in file path address
		// (Ex: filepath - "C:/users/risha/base.layout.tmpl"), for this the function gives base.layout.tmpl
		name := filepath.Base(page)

		// create a new template with the required name and the page content
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// gets all the layout templates from the templates directory
		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {

			// add the layout file to the respective template that was created above
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		//store the final template in the map with the extracted name as the key
		myCache[name] = ts
	}
	return myCache, nil
}
