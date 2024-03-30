# Adobe User Management API(UMAPI) - Go SDK
This library wraps the Adobe User Management API(UMAPI) for Go. It provides a simple way to interact with the API.

## Installation
```bash
go get github.com/bear-san/adobe-umapi
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/bear-san/adobe-umapi/pkg/api"
	"github.com/bear-san/adobe-umapi/pkg/auth"
	"github.com/bear-san/adobe-umapi/pkg/user"
	"os"
)

func main() {
	orgId := os.Getenv("ORG_ID")
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	err := auth.Setup(clientId, clientSecret)
	if err != nil {
		panic(err)
	}

	// Create a new user by Enterprise ID
	userRequests := []user.Request{
		{
			User:       "tools-test-enterprise-id@example.com",
			RequestID:  "Create_Enterprise_ID",
			Domain:     nil,
			UseAdobeID: false,
			Do: []user.Command{
				{
					CreateEnterpriseID: &user.CreateEnterpriseIDRequest{
						Email:     "tools-test-enterprise-id@example.com",
						Country:   "JP",
						FirstName: "Test",
						LastName:  "EnterpriseID",
						Option:    "updateIfAlreadyExists",
					},
				},
			},
		},
	}

	result, err := api.Exec(&userRequests, auth.Credential, orgId, clientId)
	fmt.Printf("completed: %d, notCompleted: %d, completedInTestMode: %d\n", result.Completed, result.NotCompleted, result.CompletedInTestMode)
}
```

result:
```text
completed: 1, notCompleted: 0, completedInTestMode: 0
```

## Capabilities
**to see the full functions list of UMAPI, please refer to the [User Management API Documentation](https://adobe-apiplatform.github.io/umapi-documentation/en/)**

### User Management Action API
#### User Action Commands
- [x] Create/Delete/Update Enterprise ID
- [x] Create/Delete/Update Federated ID
- [x] Add/Remove/Update Adobe ID
- [x] Add/Remove User to/from User Group
#### User Group Action Commands
- [ ] Create/Delete/Update User Group

### User Access APIs
- [ ] Get User Information
- [ ] Get Users in Organization
- [ ] Get Users by Group

## License
see [LICENSE](LICENSE)