// Package deals covers the Deals API which has been exposed to allow for easy integration with the HubSpot CRM objects.
package deals

import (
	"fmt"
	"net/http"

	"github.com/retgits/hubspot/client"
)

const (
	defaultCount  int64 = 100
	defaultOffSet int64 = 0
	// updateDealEndpoint is the endpoint to update the deals object
	updateDealEndpoint = "deals/v1/deal/%s"
	// getDealEndpoint is the endpoint to get the deals object
	getDealEndpoint = "deals/v1/deal/%s"
	// getRecentDealsEndpoint is the endpoint to get recently modified deals
	getRecentDealsEndpoint = "deals/v1/deal/recent/modified"
)

// Deals contains the elements to communicate with the HubSpot Deals endpoints.
type Deals struct {
	*client.Client
	Count  int64
	OffSet int64
}

// New creates a new instance of the Deals service with default settings.
func New(c *client.Client) *Deals {
	return &Deals{
		c, defaultCount, defaultOffSet,
	}
}

// WithOffSet sets an offset value returning a Deals pointer for
// chaining.
func (d *Deals) WithOffSet(offset int64) *Deals {
	d.OffSet = offset
	return d
}

// GetDeal returns an object representing the deal with the id :dealId associated with the specified account.
func (d *Deals) GetDeal(dealID string) (Deal, error) {
	url := buildURL(d, fmt.Sprintf(getDealEndpoint, dealID))

	res, err := d.Call(url, http.MethodGet, nil)
	if err != nil {
		return Deal{}, err
	}

	return unmarshalDeal(res)
}

// GetRecentlyModifiedDeals gets recently modified deals in an account sorted by their last modified date,
// starting with the most recently modified deals.
func (d *Deals) GetRecentlyModifiedDeals() ([]Result, error) {
	url := buildURL(d, getRecentDealsEndpoint)

	res, err := d.Call(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	temp, err := unmarshalRecentDeals(res)
	if err != nil {
		return nil, err
	}

	d = d.WithOffSet(temp.Offset)
	recentDeals := temp.Results

	if temp.HasMore {
		for {
			url := buildURL(d, getRecentDealsEndpoint)
			res, err := d.Call(url, http.MethodGet, nil)
			if err != nil {
				return nil, err
			}

			temp, err := unmarshalRecentDeals(res)
			if err != nil {
				return nil, err
			}

			d = d.WithOffSet(temp.Offset)

			recentDeals = append(recentDeals, temp.Results...)
			if !temp.HasMore {
				break
			}
		}
	}

	return recentDeals, nil
}

// UpdateDeal is to update an existing deal in HubSpot. This method lets you update the properties of a deal in HubSpot.
// The map[string]string represents the new values for the contact, where the map key is the name of the property and the map
// value is the new value
func (d *Deals) UpdateDeal(dealID string, props map[string]string) error {
	url := buildURL(d, fmt.Sprintf(updateDealEndpoint, dealID))

	properties := make([]Property, 0)

	for key, val := range props {
		prop := Property{
			Name:  key,
			Value: val,
		}
		properties = append(properties, prop)
	}

	updateProperties := Properties{
		Properties: properties,
	}

	payload, err := updateProperties.Marshal()
	if err != nil {
		return err
	}

	_, err = d.Call(url, http.MethodPut, payload)
	if err != nil {
		return err
	}

	return nil
}

// Construct the proper URL to call
func buildURL(d *Deals, u string) string {
	url := ""
	url = fmt.Sprintf("%s?hapikey=%s", u, d.APIKey)

	if d.OffSet > 0 {
		url = fmt.Sprintf("%s&offset=%d", url, d.OffSet)
	}

	return url
}
