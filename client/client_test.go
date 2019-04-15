package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	apikey = "demo" // https://developers.hubspot.com/docs/methods/auth/oauth-overview
)

func TestClient(t *testing.T) {
	hubspot := NewClient()
	assert.Equal(t, hubspot.APIKey, "")

	hubspot = NewClient().WithAPIKey(apikey)
	assert.Equal(t, hubspot.APIKey, apikey)
}
