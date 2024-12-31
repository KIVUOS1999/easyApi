package app

import (
	"net/http"

	"github.com/KIVUOS1999/easyApi/configs"
	"github.com/KIVUOS1999/easyApi/constants"
	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/gorilla/mux"
)

type handlerFunc func(ctx *Context) (interface{}, error)

type app struct {
	Muxx    *mux.Router
	Configs *configs.Config
}

func New(args ...any) *app {
	configPath := "./configs/.env"
	if len(args) > 0 {
		configPath = args[0].(string)
	}

	m := mux.NewRouter()
	config := configs.New(configPath)

	app := app{
		Muxx:    m,
		Configs: config,
	}

	return &app
}

func (a *app) Start() {
	port := a.Configs.Get(constants.Address)

	log.Info("Starting server:", port)

	err := http.ListenAndServe(port, a.Muxx)
	if err != nil {
		log.Error(err.Error())
	}
}

func (a *app) registerRoutes(path string, handler handlerFunc, method ...string) {
	a.Muxx.HandleFunc(path, a.adapterFunc(handler)).Methods(method...)
}

func (a *app) Get(path string, handler handlerFunc) {
	a.registerRoutes(path, handler, http.MethodGet, http.MethodOptions)
}

func (a *app) Post(path string, handler handlerFunc) {
	a.registerRoutes(path, handler, http.MethodPost, http.MethodOptions)
}

func (a *app) Put(path string, handler handlerFunc) {
	a.registerRoutes(path, handler, http.MethodPut, http.MethodOptions)
}

func (a *app) Delete(path string, handler handlerFunc) {
	a.registerRoutes(path, handler, http.MethodDelete, http.MethodOptions)
}

func (a *app) Options(path string, handler handlerFunc) {
	a.registerRoutes(path, handler, http.MethodOptions)
}
