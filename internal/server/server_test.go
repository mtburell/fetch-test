package server

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	s := New("localhost:8080", &mux.Router{})
	s.Setup()
	assert.NotNil(t, s.Negroni)
}
