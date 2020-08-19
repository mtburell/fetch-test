package application

import (
	"fetch-test/internal/configuration"
	"fetch-test/internal/docs"
	"fetch-test/internal/emailchecker"
	"fetch-test/internal/healthcheck"
	"fetch-test/internal/server"

	"github.com/gorilla/mux"
)

type App struct {
	Server *server.Server
}

func New(config *configuration.Config) (*App, error) {
	router := mux.NewRouter()

	healthcheck.RegisterRoutes(router)
	emailchecker.RegisterRoutes(router)
	docs.RegisterRoutes(router)

	srv := server.New(
		config.URL,
		router,
	)

	srv.Setup()

	return &App{
		Server: srv,
	}, nil

}

func (a *App) Run() {
	a.Server.Start()
}
