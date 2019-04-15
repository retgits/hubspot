// Package crmassociations covers the CRM Associations API is used to manage associations between objects in the HubSpot CRM.
// This includes the association between a contact and its company, between a company and a parent or child company, between
// deal and a company or contact, or between a ticket and a contact or company, as well as associations between engagements
// and other objects.
package crmassociations

import "encoding/json"

// Associations is a struct generated from the HubSpot API
type Associations struct {
	Results []int64 `json:"results"`
	HasMore bool    `json:"hasMore"`
	Offset  int64   `json:"offset"`
}

func unmarshalAssociations(data []byte) (Associations, error) {
	var r Associations
	err := json.Unmarshal(data, &r)
	return r, err
}
