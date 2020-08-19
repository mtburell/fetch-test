package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "", nil)
	HealthcheckHandler(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code, "Should be 200")
}
