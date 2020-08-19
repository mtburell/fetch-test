package emailchecker

import (
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
