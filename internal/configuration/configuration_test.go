package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	_, err := GetConfig()

	assert.Nil(t, err, "error should be nil")
}
