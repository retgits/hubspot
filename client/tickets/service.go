// Package tickets is part of HubSpots preview program, and should be considered as a non-stable release that
// will be subject to bugs and breaking changes while under development.
package tickets

import (
	"fmt"
	"net/http"

	"github.com/retgits/hubspot/client"
)

const (
	defaultOffSet int64 = 0
	// allTicketsEndpoint is the endpoint to retrieve all tickets
	allTicketsEndpoint = "crm-objects/v1/objects/tickets/paged"
)

// Tickets contains the elements to communicate with the HubSpot Tickets endpoints.
type Tickets struct {
	*client.Client
	OffSet     int64
	Properties []string
}

// New creates a new instance of the Tickets service with default settings.
func New(c *client.Client) *Tickets {
	return &Tickets{
		c, defaultOffSet, nil,
	}
}

// WithOffSet sets an offset value returning a Tickets pointer for
// chaining.
func (t *Tickets) WithOffSet(offset int64) *Tickets {
	t.OffSet = offset
	return t
}

// WithProperties sets the "property" parameter, then the properties in the "ticket" object in the
// returned data will only include the property or properties that you request. returning a Tickets
// pointer for chaining.
func (t *Tickets) WithProperties(props []string) *Tickets {
	t.Properties = props
	return t
}

// GetAllTickets gets all tickets from a portal. By default you will only get a few system fields for
// any tickets in the response. If you want to get specific properties, you'll need to use the properties
// parameter.
func (t *Tickets) GetAllTickets() ([]Object, error) {
	url := buildURL(t, allTicketsEndpoint)

	res, err := t.Call(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	temp, err := unmarshalTickets(res)
	if err != nil {
		return nil, err
	}

	t = t.WithOffSet(temp.Offset)
	tickets := temp.Objects

	if temp.HasMore {
		for {
			url := buildURL(t, allTicketsEndpoint)
			res, err := t.Call(url, http.MethodGet, nil)
			if err != nil {
				return nil, err
			}

			temp, err := unmarshalTickets(res)
			if err != nil {
				return nil, err
			}

			t = t.WithOffSet(temp.Offset)

			tickets = append(tickets, temp.Objects...)
			if !temp.HasMore {
				break
			}
		}
	}

	return tickets, nil
}

// Construct the proper URL to call
func buildURL(t *Tickets, u string) string {
	url := ""
	url = fmt.Sprintf("%s?hapikey=%s", u, t.APIKey)

	if t.OffSet > 0 {
		url = fmt.Sprintf("%s&offset=%d", url, t.OffSet)
	}

	for idx := range t.Properties {
		url = fmt.Sprintf("%s&properties=%s", url, t.Properties[idx])
	}

	return url
}
