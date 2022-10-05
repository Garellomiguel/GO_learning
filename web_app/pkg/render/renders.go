package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"web_app/pkg/config"
	"web_app/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Get the template cache from the app config
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// Get requested templated from cache
	t, inMap := tc[tmpl]
	if !inMap {
		log.Fatal("Template not in map")
	}
	// Render thetemplate

	td = AddDefaultData(td)
	err := t.Execute(w, td)
	if err != nil {
		log.Println("Error rendering the tempalte", err)
	}
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	match, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		if len(match) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

// func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
// 	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parseTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("Error while parsing template: ", err)
// 		return
// 	}
// }

// Another Way

// var tc = make(map[string]*template.Template) // template cache

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we alredy have the template in our cache (tc)
// 	_, inMap := tc[t]
// 	if !inMap {
// 		log.Println("Creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		fmt.Println("Template alredy created and in cache")
// 	}
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Panic("Error:", err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}
// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return nil
// 	}
// 	// add template to cache
// 	tc[t] = tmpl
// 	return nil
// }
