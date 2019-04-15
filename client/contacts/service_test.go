// Package contacts are for the fundamental building block to HubSpot, contacts. They store lead-specific data that makes it possible to leverage much of the functionality in HubSpot, from marketing automation, to lead scoring to smart content.
package contacts

import (
	"testing"

	"github.com/retgits/hubspot/client"
	"github.com/stretchr/testify/assert"
)

const (
	apikey = "demo" // https://developers.hubspot.com/docs/methods/auth/oauth-overview
)

func TestClient(t *testing.T) {
	hubspot := client.NewClient().WithAPIKey(apikey)
	contactsSvc := New(hubspot)
	assert.Equal(t, contactsSvc.APIKey, apikey)
}
