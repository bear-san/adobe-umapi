package user

type AddAdobeID struct {
	Email     string `json:"email"`
	Country   string `json:"country"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Option    string `json:"option"`
}

type CreateEnterpriseID struct {
	Email     string `json:"email"`
	Country   string `json:"country"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Option    string `json:"option"`
}

type CreateFederatedID struct {
	Email     string `json:"email"`
	Country   string `json:"country"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Option    string `json:"option"`
}
