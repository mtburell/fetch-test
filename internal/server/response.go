package server

import (
	"log"
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"error_message"`
}

func WriteResponse(res http.ResponseWriter, payload []byte) {
	res.Header().Set("Content-Type", "application/json")

	if _, err := res.Write(payload); err != nil {
		log.Panicln("Unable to write response", err)
	}
}

func WriteErrorResponse(res http.ResponseWriter, code int, message string) {
	data := &ErrorResponse{
		Code:    code,
		Message: message,
	}

	rbody, err := ffjson.Marshal(&data)
	if err != nil {
		log.Panicln("Unable to build http response", err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)

	if _, err := res.Write(rbody); err != nil {
		log.Panicln("Unable to write response", err)
	}

}
