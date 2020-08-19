package application

import (
	"testing"

	"fetch-test/internal/configuration"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	config, _ := configuration.GetConfig()
	app, err := New(config)
	assert.Nil(t, err)
	assert.NotNil(t, app)
}
