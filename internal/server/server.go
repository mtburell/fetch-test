package server

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Server struct {
	Host         string
	Negroni      *negroni.Negroni
	Router       *mux.Router
	shutdownChan chan os.Signal
}

var serverWG sync.WaitGroup

func New(hostname string, router *mux.Router) *Server {
	server := &Server{
		Host:   hostname,
		Router: router,
	}

	return server
}

func (s *Server) Setup() {
	c := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"HEAD", "OPTIONS", "POST", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{}),
		handlers.MaxAge(600),
	)

	s.Negroni = negroni.New()

	s.Negroni.UseHandler(c(s.Router))
}

func (s *Server) Start() {
	var srv http.Server

	fmt.Println("Server Running On", s.Host)
	srv.Addr = s.Host
	srv.Handler = s.Negroni
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Print("HTTPServerListenAndServeErr", err)
	}
}
