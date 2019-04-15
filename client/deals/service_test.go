// Package deals covers the Deals API which has been exposed to allow for easy integration with the HubSpot CRM objects.
package deals

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
	dealSvc := New(hubspot)
	assert.Equal(t, dealSvc.APIKey, apikey)

	newProps := make(map[string]string)
	newProps["dealname"] = "HelloWorld"

	err := dealSvc.UpdateDeal("680305641", newProps)
	assert.NoError(t, err)
}
