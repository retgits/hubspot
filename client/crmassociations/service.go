// Package crmassociations covers the CRM Associations API is used to manage associations between objects in the HubSpot CRM.
// This includes the association between a contact and its company, between a company and a parent or child company, between
// deal and a company or contact, or between a ticket and a contact or company, as well as associations between engagements
// and other objects.
package crmassociations

import (
	"fmt"
	"net/http"

	"github.com/retgits/hubspot/client"
)

const (
	defaultOffSet int64 = 0
	defaultLimit  int64 = 100
	// associationsEndpoint is the endpoint to retrieve all associations for an object
	associationsEndpoint = "crm-associations/v1/associations/%s/HUBSPOT_DEFINED/%s"
)

// CRMAssociations contains the elements to communicate with the HubSpot CRM Associations endpoints.
type CRMAssociations struct {
	*client.Client
	OffSet int64
	Limit  int64
}

// New creates a new instance of the CRMAssociations service with default settings.
func New(c *client.Client) *CRMAssociations {
	return &CRMAssociations{
		c, defaultOffSet, defaultLimit,
	}
}

// WithOffSet sets an offset value returning a CRMAssociations pointer for
// chaining.
func (c *CRMAssociations) WithOffSet(offset int64) *CRMAssociations {
	c.OffSet = offset
	return c
}

// WithLimit sets a limit value returning a CRMAssociations pointer for
// chaining.
func (c *CRMAssociations) WithLimit(limit int64) *CRMAssociations {
	c.Limit = limit
	return c
}

// GetAssociationsForCRMObject gets the IDs of objects associated with the given object, based on the specified association type.
func (c *CRMAssociations) GetAssociationsForCRMObject(objectID string, definitionID string) ([]int64, error) {
	url := buildURL(c, fmt.Sprintf(associationsEndpoint, objectID, definitionID))

	res, err := c.Call(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	temp, err := unmarshalAssociations(res)
	if err != nil {
		return nil, err
	}

	c = c.WithOffSet(temp.Offset)
	associations := temp.Results

	if temp.HasMore {
		for {
			url := buildURL(c, fmt.Sprintf(associationsEndpoint, objectID, definitionID))
			res, err := c.Call(url, http.MethodGet, nil)
			if err != nil {
				return nil, err
			}

			temp, err := unmarshalAssociations(res)
			if err != nil {
				return nil, err
			}

			c = c.WithOffSet(temp.Offset)

			associations = append(associations, temp.Results...)
			if !temp.HasMore {
				break
			}
		}
	}

	return associations, nil
}

// Construct the proper URL to call
func buildURL(c *CRMAssociations, u string) string {
	url := ""
	url = fmt.Sprintf("%s?hapikey=%s", u, c.APIKey)

	if c.OffSet > 0 {
		url = fmt.Sprintf("%s&offset=%d", url, c.OffSet)
	}

	if c.Limit > 0 {
		url = fmt.Sprintf("%s&limit=%d", url, c.Limit)
	}

	return url
}
