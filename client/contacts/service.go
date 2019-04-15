// Package contacts are for the fundamental building block to HubSpot, contacts. They store lead-specific data that makes it possible to leverage much of the functionality in HubSpot, from marketing automation, to lead scoring to smart content.
package contacts

import (
	"fmt"
	"net/http"

	"github.com/retgits/hubspot/client"
)

const (
	defaultCount  int64 = 100
	defaultOffSet int64 = 0
	// recentlyUpdatedcontactsEndpoint is the endpoint to retrieve the recently updated and created contacts
	recentlyUpdatedcontactsEndpoint = "contacts/v1/lists/recently_updated/contacts/recent"
	// updateContactsEndpoint is the endpoint to update contacts
	updateContactsEndpoint = "contacts/v1/contact/vid/%s/profile"
)

// Contacts contains the elements to communicate with the HubSpot Contacts endpoints.
type Contacts struct {
	*client.Client
	Count      int64
	OffSet     int64
	Properties []string
}

// New creates a new instance of the Contacts service with default settings. Because of the central role contacts play in the HubSpot application,
// it is not surprising that most integrations with HubSpot either read or write Contacts data.
func New(c *client.Client) *Contacts {
	return &Contacts{
		c, defaultCount, defaultOffSet, nil,
	}
}

// WithCount sets an account value returning a Contacts pointer for
// chaining.
func (c *Contacts) WithCount(count int64) *Contacts {
	c.Count = count
	return c
}

// WithOffSet sets an offset value returning a Contacts pointer for
// chaining.
func (c *Contacts) WithOffSet(offset int64) *Contacts {
	c.OffSet = offset
	return c
}

// WithProperties sets the "property" parameter, then the properties in the "contact" object in the
// returned data will only include the property or properties that you request. returning a Contacts
// pointer for chaining.
func (c *Contacts) WithProperties(props []string) *Contacts {
	c.Properties = props
	return c
}

// GetRecentlyUpdatedContacts returns, for a given account, all contacts that have been recently updated or created.
func (c *Contacts) GetRecentlyUpdatedContacts() ([]Contact, error) {
	url := url(c, recentlyUpdatedcontactsEndpoint)

	res, err := c.Call(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	temp, err := unmarshalHubSpotContacts(res)
	if err != nil {
		return nil, err
	}

	c = c.WithOffSet(temp.VidOffset)
	pax := temp.Contacts

	if temp.HasMore {
		for {
			res, err := c.Call(url, http.MethodGet, nil)
			if err != nil {
				return nil, err
			}

			temp, err := unmarshalHubSpotContacts(res)
			if err != nil {
				return nil, err
			}

			c = c.WithOffSet(temp.VidOffset)

			pax = append(pax, temp.Contacts...)
			if !temp.HasMore {
				break
			}
		}
	}

	return pax, nil
}

// UpdateContact is to update an existing contact in HubSpot. This method lets you update the properties of a contact in HubSpot.
// The map[string]string represents the new values for the contact, where the map key is the name of the property and the map
// value is the new value
func (c *Contacts) UpdateContact(contactID string, props map[string]string) error {
	url := url(c, fmt.Sprintf(updateContactsEndpoint, contactID))

	properties := make([]Property, 0)

	for key, val := range props {
		prop := Property{
			Property: key,
			Value:    val,
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

	_, err = c.Call(url, http.MethodPost, payload)
	if err != nil {
		return err
	}

	return nil
}

// Construct the proper URL to call
func url(c *Contacts, u string) string {
	url := ""
	url = fmt.Sprintf("%s?hapikey=%s", u, c.APIKey)

	if c.OffSet > 0 {
		url = fmt.Sprintf("%s&vidOffset=%d", url, c.OffSet)
	}

	if c.Count > 0 {
		url = fmt.Sprintf("%s&count=%d", url, c.Count)
	}

	for idx := range c.Properties {
		url = fmt.Sprintf("%s&property=%s", url, c.Properties[idx])
	}

	return url
}
