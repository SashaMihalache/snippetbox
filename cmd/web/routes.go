package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes(cfg config) http.Handler {
	// mux := http.NewServeMux()

	// fileServer := http.FileServer(http.Dir(cfg.staticDir))
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/snippet/view", app.snippetView)
	// mux.HandleFunc("/snippet/create", app.snippetCreate)

	// standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// return standard.Then(mux)

	router := httprouter.New()

	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
