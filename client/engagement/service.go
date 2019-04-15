// Package engagement covers the engagements which are used to store data from CRM actions,
// including notes, tasks, meetings, and calls.
package engagement

import (
	"fmt"
	"net/http"

	"github.com/retgits/hubspot/client"
)

const (
	// engagementEndpoint is the endpoint to retrieve a single engagement
	engagementEndpoint = "engagements/v1/engagements/%d"
)

// Engagements contains the elements to communicate with the HubSpot Engagements endpoints.
type Engagements struct {
	*client.Client
}

// New creates a new instance of the Engagements service with default settings.
func New(c *client.Client) *Engagements {
	return &Engagements{
		c,
	}
}

// GetEngagement gets an engagement (a task or activity) on an object in HubSpot.
func (e *Engagements) GetEngagement(engagementID int64) (HubspotEngagement, error) {
	url := url(e, fmt.Sprintf(engagementEndpoint, engagementID))

	res, err := e.Call(url, http.MethodGet, nil)
	if err != nil {
		return HubspotEngagement{}, err
	}

	return unmarshalHubspotEngagement(res)
}

// Construct the proper URL to call
func url(e *Engagements, u string) string {
	url := ""
	url = fmt.Sprintf("%s?hapikey=%s", u, e.APIKey)

	return url
}
