package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"callback"
	"home"
	"login"
	"logout"
	"middlewares"
	"user"
)

// StartServer whatever
func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.Handler)
	r.HandleFunc("/login", login.Handler)
	r.HandleFunc("/logout", logout.Handler)
	r.HandleFunc("/callback", callback.Handler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.Handler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	log.Print("Server listening on http://localhost:3000/")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
