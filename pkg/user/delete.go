package user

type RemoveFromGroupRequest struct {
	Group []string `json:"group"`
}

type RemoveFromAllGroupsRequest struct {
	Group string `json:"group"`
}

type RemoveFromOrgRequest struct {
	DeleteAccount bool `json:"deleteAccount"`
}
