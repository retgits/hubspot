// Package engagement covers the engagements which are used to store data from CRM actions,
// including notes, tasks, meetings, and calls.
package engagement

import "encoding/json"

// HubspotEngagement is a struct generated from the HubSpot API
type HubspotEngagement struct {
	Engagement   Engagement    `json:"engagement"`
	Associations Associations  `json:"associations"`
	Attachments  []interface{} `json:"attachments"`
	Metadata     Metadata      `json:"metadata"`
}

// Associations is a struct generated from the HubSpot API
type Associations struct {
	ContactIDS  []int64       `json:"contactIds"`
	CompanyIDS  []interface{} `json:"companyIds"`
	DealIDS     []interface{} `json:"dealIds"`
	OwnerIDS    []interface{} `json:"ownerIds"`
	WorkflowIDS []interface{} `json:"workflowIds"`
	TicketIDS   []interface{} `json:"ticketIds"`
	ContentIDS  []interface{} `json:"contentIds"`
	QuoteIDS    []interface{} `json:"quoteIds"`
}

// Engagement is a struct generated from the HubSpot API
type Engagement struct {
	ID                   int64         `json:"id"`
	PortalID             int64         `json:"portalId"`
	Active               bool          `json:"active"`
	CreatedAt            int64         `json:"createdAt"`
	LastUpdated          int64         `json:"lastUpdated"`
	CreatedBy            int64         `json:"createdBy"`
	ModifiedBy           int64         `json:"modifiedBy"`
	OwnerID              int64         `json:"ownerId"`
	Type                 string        `json:"type"`
	Timestamp            int64         `json:"timestamp"`
	AllAccessibleTeamIDS []interface{} `json:"allAccessibleTeamIds"`
	BodyPreview          string        `json:"bodyPreview"`
	QueueMembershipIDS   []interface{} `json:"queueMembershipIds"`
}

// Metadata is a struct generated from the HubSpot API
type Metadata struct {
	Body string `json:"body"`
}

func unmarshalHubspotEngagement(data []byte) (HubspotEngagement, error) {
	var r HubspotEngagement
	err := json.Unmarshal(data, &r)
	return r, err
}
