// Package contacts are for the fundamental building block to HubSpot, contacts. They store lead-specific data that makes it possible to leverage much of the functionality in HubSpot, from marketing automation, to lead scoring to smart content.
package contacts

import "encoding/json"

// HubSpotContacts is the payload returned after calling the Contacts API
type HubSpotContacts struct {
	Contacts   []Contact `json:"contacts"`
	HasMore    bool      `json:"has-more"`
	VidOffset  int64     `json:"vid-offset"`
	TimeOffset int64     `json:"time-offset"`
}

// Contact is a single contact in HubSpot
type Contact struct {
	AddedAt          int64                        `json:"addedAt"`
	Vid              int64                        `json:"vid"`
	CanonicalVid     int64                        `json:"canonical-vid"`
	MergedVids       []interface{}                `json:"merged-vids"`
	PortalID         int64                        `json:"portal-id"`
	IsContact        bool                         `json:"is-contact"`
	ProfileToken     string                       `json:"profile-token"`
	ProfileURL       string                       `json:"profile-url"`
	Properties       map[string]map[string]string `json:"properties"`
	FormSubmissions  []interface{}                `json:"form-submissions"`
	IdentityProfiles []IdentityProfile            `json:"identity-profiles"`
	MergeAudits      []interface{}                `json:"merge-audits"`
}

// IdentityProfile is a struct generated from the HubSpot API
type IdentityProfile struct {
	Vid                     int64      `json:"vid"`
	SavedAtTimestamp        int64      `json:"saved-at-timestamp"`
	DeletedChangedTimestamp int64      `json:"deleted-changed-timestamp"`
	Identities              []Identity `json:"identities"`
}

// Identity is a struct generated from the HubSpot API
type Identity struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
	IsPrimary *bool  `json:"is-primary,omitempty"`
}

// Properties is a struct generated from the HubSpot API
type Properties struct {
	Properties []Property `json:"properties"`
}

// Property is a struct generated from the HubSpot API
type Property struct {
	Property string `json:"property,omitempty"`
	Value    string `json:"value"`
}

// Marshal takes an Properties struct and transforms it into a byte array
func (r *Properties) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func unmarshalHubSpotContacts(data []byte) (HubSpotContacts, error) {
	var r HubSpotContacts
	err := json.Unmarshal(data, &r)
	return r, err
}
