package user

type AddAdobeIDRequest struct {
	Email     string `json:"email"`
	Country   string `json:"country"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Option    string `json:"option"`
}

type CreateEnterpriseIDRequest struct {
	Email     string `json:"email"`
	Country   string `json:"country"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Option    string `json:"option"`
}

type CreateFederatedIDRequest struct {
	Email     string `json:"email"`
	Country   string `json:"country"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Option    string `json:"option"`
}
