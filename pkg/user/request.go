package user

type Request struct {
	User       string    `json:"user"`
	RequestID  string    `json:"requestID"`
	Domain     string    `json:"domain"`
	UseAdobeID bool      `json:"useAdobeID"`
	Do         []Command `json:"do"`
}

type Command struct {
	AddAdobeID          *AddAdobeIDRequest          `json:"addAdobeID,omitempty"`
	CreateEnterpriseID  *CreateEnterpriseIDRequest  `json:"createEnterpriseID,omitempty"`
	CreateFederatedID   *CreateFederatedIDRequest   `json:"createFederatedID,omitempty"`
	AddGroup            *AddGroupRequest            `json:"addGroup,omitempty"`
	Update              *UpdateRequest              `json:"update,omitempty"`
	RemoveFromGroup     *RemoveFromGroupRequest     `json:"removeFromGroup,omitempty"`
	RemoveFromAllGroups *RemoveFromAllGroupsRequest `json:"removeFromAllGroups,omitempty"`
	RemoveFromOrg       *RemoveFromOrgRequest       `json:"removeFromOrg,omitempty"`
}
