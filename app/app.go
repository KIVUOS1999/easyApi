package app

import (
	"net/http"

	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/gorilla/mux"
)

type handlerFunc func(ctx *Context) (interface{}, error)

type app struct {
	muxx *mux.Router
}

func New() *app {
	m := mux.NewRouter()

	app := app{
		muxx: m,
	}

	return &app
}

func (a *app) Start() {
	port := ":8000"

	log.Info("Starting server:", port)

	http.ListenAndServe(port, a.muxx)
}

func (a *app) registerRoutes(path string, method string, handler handlerFunc) {
	a.muxx.HandleFunc(path, a.adapterFunc(handler)).Methods(method)
}

func (a *app) Get(path string, handler handlerFunc) {
	a.registerRoutes(path, http.MethodGet, handler)
}

func (a *app) Post(path string, handler handlerFunc) {
	a.registerRoutes(path, http.MethodPost, handler)
}

func (a *app) Put(path string, handler handlerFunc) {
	a.registerRoutes(path, http.MethodPut, handler)
}

func (a *app) Delete(path string, handler handlerFunc) {
	a.registerRoutes(path, http.MethodDelete, handler)
}
