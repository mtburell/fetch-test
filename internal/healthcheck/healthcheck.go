package healthcheck

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	router.Path("/healthcheck").
		HandlerFunc(HealthcheckHandler).
		Methods("GET")
}

func HealthcheckHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
}
