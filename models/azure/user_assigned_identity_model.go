package azure

type UserAssignedIdentityModel struct {
	Entity

	Location   string               `json:"location,omitempty"`
	Name       string               `json:"name,omitempty"`
	Properties UserAssignedIdentity `json:"properties,omitempty"`
	Type       string               `json:"type,omitempty"`
	Tags       map[string]string    `json:"tags,omitempty"`
}
