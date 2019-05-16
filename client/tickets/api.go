// Package tickets is part of HubSpots preview program, and should be considered as a non-stable release that
// will be subject to bugs and breaking changes while under development.
package tickets

import (
	"encoding/json"
)

type HubSpotTickets struct {
	Objects []Object `json:"objects"`
	HasMore bool     `json:"hasMore"`
	Offset  int64    `json:"offset"`
}

type Object struct {
	ObjectType string                            `json:"objectType"`
	PortalID   int64                             `json:"portalId"`
	ObjectID   int64                             `json:"objectId"`
	Properties map[string]map[string]interface{} `json:"properties"`
	Version    int64                             `json:"version"`
	IsDeleted  bool                              `json:"isDeleted"`
}

func unmarshalTickets(data []byte) (HubSpotTickets, error) {
	var r HubSpotTickets
	err := json.Unmarshal(data, &r)
	return r, err
}
