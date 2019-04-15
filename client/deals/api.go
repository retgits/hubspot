// Package deals covers the Deals API which has been exposed to allow for easy integration with the HubSpot CRM objects.
package deals

import "encoding/json"

// Associations is a struct generated from the HubSpot API
type Associations struct {
	AssociatedVids       []int64       `json:"associatedVids"`
	AssociatedCompanyIDS []interface{} `json:"associatedCompanyIds"`
	AssociatedDealIDS    []interface{} `json:"associatedDealIds"`
	AssociatedTicketIDS  []interface{} `json:"associatedTicketIds"`
}

// Deal is a struct generated from the HubSpot API
type Deal struct {
	PortalID     int64               `json:"portalId"`
	DealID       int64               `json:"dealId"`
	IsDeleted    bool                `json:"isDeleted"`
	Associations Associations        `json:"associations"`
	Properties   map[string]Property `json:"properties"`
	Imports      []interface{}       `json:"imports"`
	StateChanges []interface{}       `json:"stateChanges"`
}

// Properties is a struct generated from the HubSpot API
type Properties struct {
	Properties []Property `json:"properties"`
}

// Property is a struct generated from the HubSpot API
type Property struct {
	Name      string    `json:"name,omitempty"`
	Value     string    `json:"value"`
	Timestamp int64     `json:"timestamp,omitempty"`
	Source    string    `json:"source,omitempty"`
	SourceID  string    `json:"sourceId,omitempty"`
	Versions  []Version `json:"versions,omitempty"`
}

// RecentDeals is a struct generated from the HubSpot API
type RecentDeals struct {
	Results []Result `json:"results"`
	HasMore bool     `json:"hasMore"`
	Offset  int64    `json:"offset"`
	Total   int64    `json:"total"`
}

// Result is a struct generated from the HubSpot API
type Result struct {
	PortalID     int64               `json:"portalId"`
	DealID       int64               `json:"dealId"`
	IsDeleted    bool                `json:"isDeleted"`
	Associations Associations        `json:"associations"`
	Properties   map[string]Property `json:"properties"`
	Imports      []interface{}       `json:"imports"`
	StateChanges []interface{}       `json:"stateChanges"`
}

// Version is a struct generated from the HubSpot API
type Version struct {
	Name      string        `json:"name"`
	Value     string        `json:"value"`
	Timestamp int64         `json:"timestamp"`
	SourceID  string        `json:"sourceId,omitempty"`
	Source    string        `json:"source"`
	SourceVid []interface{} `json:"sourceVid"`
}

// Marshal takes an Properties struct and transforms it into a byte array
func (r *Properties) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func unmarshalDeal(data []byte) (Deal, error) {
	var r Deal
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalRecentDeals(data []byte) (RecentDeals, error) {
	var r RecentDeals
	err := json.Unmarshal(data, &r)
	return r, err
}
