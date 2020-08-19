package docs

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/docs").Handler(
		negroni.New(
			&negroni.Static{
				Dir:       http.Dir("public"),
				Prefix:    "/docs",
				IndexFile: "index.html",
			},
		))
}
