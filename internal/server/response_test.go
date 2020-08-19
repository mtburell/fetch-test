package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	data := make(map[string]interface{})
	data["testing"] = "This is at test."
	body, _ := json.Marshal(&data)
	WriteResponse(recorder, body)
	assert.Equal(t, http.StatusOK, recorder.Code, "Response code should be 200")

}

func TestWriteErrorResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	WriteErrorResponse(recorder, 500, "This software is badly broken. The creator must be truly unaware")
	assert.Equal(t, http.StatusInternalServerError, recorder.Code, "The response codes should match")
}
