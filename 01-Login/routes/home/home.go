package home

import (
	"net/http"

	"templates"
)

// Handler whatever
func Handler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}
