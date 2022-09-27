package handlers

import (
	"net/http"

	"github.com/nomansalhab/golang_training_project/pkg/render"
)

// Home Page Handler
func Home(w http.ResponseWriter, r *http.Request) {
	// _, _ = fmt.Fprintf(w, "This is The Home Page")
	render.RenderTemplate(w, "home.page.html")
}

// About Page Handler
func About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is The About Page and 2 + 2 is: %d", sum))
	render.RenderTemplate(w, "about.page.html")
}
