package emailchecker

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailValid(t *testing.T) {
	// Good examples
	result := isEmailValid("test@gmail.com")
	assert.True(t, result, "Should be true")

	result = isEmailValid("test.test@gmail.com")
	assert.True(t, result, "Should be true")

	result = isEmailValid("test+test@gmail.com")
	assert.True(t, result, "Should be true")

	// Bad examples
	result = isEmailValid("test")
	assert.False(t, result, "Should be false")

	result = isEmailValid("test.com")
	assert.False(t, result, "Should be false")

	result = isEmailValid("test@test@gmail.com")
	assert.False(t, result, "Should be false")
}

func TestValidateTargets(t *testing.T) {
	emails := make([]string, 3)
	emails = append(emails, "test.email@gmail.com")
	emails = append(emails, "test.email+spam@gmail.com")
	emails = append(emails, "testemail@gmail.com")

	targets := make(map[string]bool)

	validateTargets(emails, targets)
	assert.Equal(t, 1, len(targets), "should be equal")
}

func TestHealthcheck(t *testing.T) {
	recorder := httptest.NewRecorder()

	var data = []byte(fmt.Sprintf(`{
		"emails": ["test@gmail.com"]
  }`))
	req, _ := http.NewRequest("POST", "/email-checker", bytes.NewBuffer(data))
	EmailCheckerHandler(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code, "Should be 200")
}

func TestHealthcheckBadRequest(t *testing.T) {
	recorder := httptest.NewRecorder()

	var data = []byte(fmt.Sprintf(`{
		"emails": ["test@gmail.com]
  }`))
	req, _ := http.NewRequest("POST", "/email-checker", bytes.NewBuffer(data))
	EmailCheckerHandler(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code, "Should be 400")
}
