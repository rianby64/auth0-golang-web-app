package user

import (
	"net/http"

	"app"
	"templates"
)

// Handler whatever
func Handler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile, _ := session.Values["profile"]
	templates.RenderTemplate(w, "user", profile)
}
