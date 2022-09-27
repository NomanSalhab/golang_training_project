package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate Function
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err1 := CreateTemplateCache()
	if err1 != nil {
		fmt.Println("CreateTemplateCache Error:", err1)
	} else {
		t, ok := tc[tmpl]
		if !ok {
			// log.Fatal(err1)
			fmt.Println("!ok Error:", err1)
		}
		buf := new(bytes.Buffer)
		err2 := t.Execute(buf, nil)
		if err2 != nil {
			fmt.Println("Error at t.Execute", err2)
		}
		_, err := buf.WriteTo(w)
		if err != nil {
			fmt.Println("Error writing template to browser", err)
		}
		// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
		// err := parsedTemplate.Execute(w, nil)
		// if err != nil {
		// 	fmt.Println("Error parsing template:", err)
		// 	return
		// }
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	} else {
		for _, page := range pages {
			name := filepath.Base(page)
			ts, err := template.New(name).Funcs(functions).ParseFiles(page)
			if err != nil {
				return myCache, err
			} else {
				matches, err := filepath.Glob("./templates/*.layout.html")
				if err != nil {
					return myCache, err
				} else {
					if len(matches) > 0 {
						ts, err = ts.ParseGlob("./templates/*.layout.html")
						if err != nil {
							return myCache, err
						} else {
							myCache[name] = ts
							return myCache, nil
						}
					}
				}
			}
		}
	}
	return myCache, nil
}
